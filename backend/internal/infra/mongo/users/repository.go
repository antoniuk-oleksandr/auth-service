package users

import (
	"context"
	"errors"

	usersDomain "github.com/antoniuk-oleksandr/auth-service/backend/internal/domain/users"
	"github.com/antoniuk-oleksandr/auth-service/backend/internal/logger"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

const collectionName = "users"

type repository struct {
	collection *mongo.Collection
	mapper     Mapper
	lgr        logger.Logger
}

func NewRepository(db *mongo.Database, mapper Mapper, lgr logger.Logger) usersDomain.Repository {
	return &repository{
		collection: db.Collection(collectionName),
		mapper:     mapper,
		lgr:        lgr,
	}
}

func (repo *repository) FindByUsername(ctx context.Context, username string) (*usersDomain.User, error) {
	var user User

	filter := bson.M{"username": username}
	err := repo.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			repo.lgr.Warn("User not found",
				logger.NewStringField("username", username),
			)
			return nil, usersDomain.ErrUserNotFound
		}

		repo.lgr.Error("Failed to find user by username",
			logger.NewStringField("username", username),
			logger.NewErrField(err),
		)
		return nil, usersDomain.ErrFailedToFindUser
	}

	repo.lgr.Info("User fetched successfully",
		logger.NewStringField("username", username),
	)

	userDomain := repo.mapper.MapUserModelToDomain(user)
	return &userDomain, nil
}

func (repo *repository) FindByID(ctx context.Context, id string) (*usersDomain.User, error) {
	var user User

	oid, err := bson.ObjectIDFromHex(id)
	if err != nil {
		repo.lgr.Error("Invalid user ID format",
			logger.NewStringField("id", id),
			logger.NewErrField(err),
		)
		return nil, err
	}

	filter := bson.M{"_id": oid}
	err = repo.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			repo.lgr.Warn("User not found by ID",
				logger.NewStringField("id", id),
			)
			return nil, usersDomain.ErrUserNotFound
		}

		repo.lgr.Error("Failed to find user by ID",
			logger.NewStringField("id", id),
			logger.NewErrField(err),
		)
		return nil, usersDomain.ErrFailedToFindUser
	}

	repo.lgr.Info("User fetched successfully by ID",
		logger.NewStringField("id", id),
	)

	userDomain := repo.mapper.MapUserModelToDomain(user)
	return &userDomain, nil
}

func (repo *repository) Create(
	ctx context.Context,
	user usersDomain.CreateUserCommand,
) (*usersDomain.User, error) {
	id := primitive.NewObjectID()

	doc := bson.D{
		{Key: "_id", Value: id},
		{Key: "username", Value: user.Username},
		{Key: "passwordHash", Value: user.PasswordHash},
	}

	_, err := repo.collection.InsertOne(ctx, doc)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			repo.lgr.Warn("Username is already taken",
				logger.NewStringField("username", user.Username),
			)
			return nil, usersDomain.ErrUsernameTaken
		}

		repo.lgr.Error("Failed to create user",
			logger.NewStringField("username", user.Username),
			logger.NewErrField(err),
		)
		return nil, usersDomain.ErrFailedToCreateUser
	}

	repo.lgr.Info("User created successfully",
		logger.NewStringField("id", id.Hex()),
		logger.NewStringField("username", user.Username),
	)

	return &usersDomain.User{
		ID:           id.Hex(),
		Username:     user.Username,
		PasswordHash: user.PasswordHash,
	}, nil
}
