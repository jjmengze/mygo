package repo

import (
	"context"
	"github.com/jjmengze/mygo/internal/model"
	"github.com/jjmengze/mygo/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	GetUser(ctx context.Context, condition model.QueryUser) (model.User, error)
	CreateUser(ctx context.Context, data *model.User) (*model.User, error)
}

// GetUser rdbms get user
func (repo *repository) GetUser(ctx context.Context, condition model.QueryUser) (model.User, error) {
	panic("implement me")
}

// CreateUser rdbms create user
func (repo *repository) CreateUser(ctx context.Context, data *model.User) (*model.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return &model.User{}, err
	}
	data.Password = string(hashedPassword)
	err = repo._writeDB.WithContext(ctx).Create(data).Error
	if err != nil {
		return &model.User{}, errors.ConvertMySQLError(err)
	}
	return data, nil
}
