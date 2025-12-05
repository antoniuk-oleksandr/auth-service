package mongodb

import (
	"context"
	"github.com/antoniuk-oleksandr/auth-service/backend/internal/db"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type txManager struct {
	client *mongo.Client
}

func NewTransactionManager(client *mongo.Client) db.TxManager {
	return &txManager{
		client: client,
	}
}

func (t *txManager) RunInTx(
	ctx context.Context,
	fn func(txCtx context.Context) (any, error),
) (any, error) {
	sess, err := t.client.StartSession()
	if err != nil {
		return nil, err
	}
	defer sess.EndSession(ctx)

	result, err := sess.WithTransaction(ctx, func(txCtx context.Context) (any, error) {
		return fn(txCtx)
	})

	return result, err
}
