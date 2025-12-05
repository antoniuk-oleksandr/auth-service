package users

import (
	usersDomain "github.com/antoniuk-oleksandr/auth-service/backend/internal/domain/users"
	"github.com/antoniuk-oleksandr/auth-service/backend/internal/logger"
	"context"
	"errors"

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
			return nil, usersDomain.ErrUserNotFound
		}
		return nil, usersDomain.ErrFailedToFindUser
	}

	userDomain := repo.mapper.MapUserModelToDomain(user)
	return &userDomain, nil
}

func (repo *repository) FindByID(ctx context.Context, id string) (*usersDomain.User, error) {
	var user User

	oid, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": oid}
	err = repo.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, usersDomain.ErrUserNotFound
		}
		return nil, usersDomain.ErrFailedToFindUser
	}

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

	if _, err := repo.collection.InsertOne(ctx, doc); err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return nil, usersDomain.ErrUsernameTaken
		}

		return nil, usersDomain.ErrFailedToCreateUser
	}

	return &usersDomain.User{
		ID:           id.Hex(),
		Username:     user.Username,
		PasswordHash: user.PasswordHash,
	}, nil
}
