package mongodb

import (
	"github.com/antoniuk-oleksandr/auth-service/backend/internal/config"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func New(cfg config.MongoConfig) (*mongo.Client, error) {
	return mongo.Connect(options.Client().ApplyURI(cfg.URI))
}
