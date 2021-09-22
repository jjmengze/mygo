package repo

import (
	"context"
	"github.com/jjmengze/mygo/internal/model"
)

type UserRepository interface {
	GetUser(ctx context.Context, condition model.QueryUser) (model.User, error)
}

// GetUser rdbms get user
func (repo *repository) GetUser(ctx context.Context, condition model.QueryUser) (model.User, error) {
	panic("implement me")
}

