package repo

import (
	"context"
	"github.com/jjmengze/mygo/internal/model"
	errMsg "github.com/jjmengze/mygo/pkg/errors"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"k8s.io/klog"
)

type UserRepository interface {
	GetUser(ctx context.Context, condition *model.QueryUser) (*model.User, error)
	CreateUser(ctx context.Context, data *model.User) (*model.User, error)
	UpdateAccountOTPSecret(ctx context.Context, user *model.User, opts *model.UpdateUserWhereOpts) error
}

// GetUser rdbms get user
func (repo *repository) GetUser(ctx context.Context, condition *model.QueryUser) (*model.User, error) {
	var user model.User
	err := repo._readDB.WithContext(ctx).
		Clauses(condition.Clause()...).
		Scopes(condition.Where, condition.Preload).
		First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.User{}, errors.Wrap(errMsg.ErrResourceNotFound, "get User from database error")
		}
		return &model.User{}, errors.Wrap(err, "get address from database error")
	}
	return &user, nil
}

// CreateUser rdbms create user
func (repo *repository) CreateUser(ctx context.Context, data *model.User) (*model.User, error) {
	err := repo._writeDB.WithContext(ctx).Create(data).Error
	if err != nil {
		return &model.User{}, errMsg.ConvertMySQLError(err)
	}
	return data, nil
}

// GenerateOTPAuth 傳入 accountID, 建立該ID的OTP UUID
func (repo *repository) UpdateAccountOTPSecret(ctx context.Context, user *model.User, opts *model.UpdateUserWhereOpts) error {
	err := repo._writeDB.WithContext(ctx).
		Clauses(opts.Clause()...).
		Scopes(opts.Where).
		Updates(&user).Error
	if err != nil {
		klog.Errorf("database: UpdateAccount : %#v fail: %v", user, err)
		return err
	}
	return nil
}
