package usecase

import (
	"context"
	"github.com/jjmengze/mygo/internal/model"
	"github.com/jjmengze/mygo/internal/repo"
)

// UserService Address service implement...
type UserService interface {
	GetUser(ctx context.Context, condition model.QueryUser) (model.User, error)
	ListUser(ctx context.Context, condition model.QueryUser) ([]model.User, error)
	CreateUser(ctx context.Context, address *model.User) (model.User, error)
	UpdateUser(ctx context.Context, updateAddress model.User, opts model.UpdateUserWhereOpts) error
	DeleteUser(ctx context.Context, condition model.QueryUser) error
}

// UserService  service  ...
type userService struct {
	repo repo.Repository
}

func (u userService) GetUser(ctx context.Context, condition model.QueryUser) (model.User, error) {
	panic("implement me")
}

func (u userService) ListUser(ctx context.Context, condition model.QueryUser) ([]model.User, error) {
	panic("implement me")
}

func (u userService) CreateUser(ctx context.Context, address *model.User) (model.User, error) {
	panic("implement me")
}

func (u userService) UpdateUser(ctx context.Context, updateAddress model.User, opts model.UpdateUserWhereOpts) error {
	panic("implement me")
}

func (u userService) DeleteUser(ctx context.Context, condition model.QueryUser) error {
	panic("implement me")
}

// NewUserService new service constructor
func NewUserService(repo repo.Repository) UserService {
	return &userService{
		repo: repo,
	}
}
