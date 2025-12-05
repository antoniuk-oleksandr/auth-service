package db

import "context"

type TxManager interface {
	RunInTx(ctx context.Context, fn func(sessCtx context.Context) (any, error)) (any, error)
}
