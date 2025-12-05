package auth

import (
	"github.com/antoniuk-oleksandr/auth-service/backend/internal/db"
	authDomain "github.com/antoniuk-oleksandr/auth-service/backend/internal/domain/auth"
	"context"
	"fmt"
)

type txService struct {
	txManager   db.TxManager
	authService authDomain.Service
}

func NewServiceWithTx(
	authService authDomain.Service,
	txManager db.TxManager,
) authDomain.Service {
	return &txService{
		txManager:   txManager,
		authService: authService,
	}
}

func (t *txService) Login(ctx context.Context, cmd authDomain.LoginCommand) (*authDomain.JWT, error) {
	return t.authService.Login(ctx, cmd)
}

func (t *txService) Register(ctx context.Context, cmd authDomain.RegisterCommand) (*authDomain.JWT, error) {
	return t.runWithTx(ctx, func(sessCtx context.Context) (*authDomain.JWT, error) {
		return t.authService.Register(sessCtx, cmd)
	})
}

func (t *txService) runWithTx(ctx context.Context, fn func(sessCtx context.Context) (*authDomain.JWT, error)) (*authDomain.JWT, error) {
	result, err := t.txManager.RunInTx(ctx, func(sessCtx context.Context) (any, error) {
		return fn(sessCtx)
	})
	if err != nil {
		return nil, err
	}

	res, ok := result.(*authDomain.JWT)
	if !ok {
		return nil, fmt.Errorf("failed to cast result to *authDomain.JWT")
	}

	return res, nil
}
