package usecase

import (
	"context"
	"github.com/google/uuid"
	"github.com/jjmengze/mygo/internal/model"
	"github.com/jjmengze/mygo/internal/repo"
	errMsg "github.com/jjmengze/mygo/pkg/errors"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(ctx context.Context, condition *model.QueryUser) (*model.User, error)
}

// UserService  service  ...
type authService struct {
	repo repo.Repository
}

// NewAuthService new service constructor
func NewAuthService(repo repo.Repository) AuthService {
	return &authService{
		repo: repo,
	}
}

func (a *authService) Login(ctx context.Context, condition *model.QueryUser) (*model.User, error) {
	//condition.User.Password
	password := condition.User.Password
	condition.User.Password = ""
	user, err := a.repo.GetUser(ctx, condition)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		//if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		//
		//}
		return nil, errors.WithMessage(errMsg.ErrUsernameOrPasswordIncorrect, err.Error())
	}
	return user, nil
}

func (a *authService) GenerateOTPAuth(ctx context.Context, opts *model.UpdateUserWhereOpts) (string, error) {

	user := &model.User{
		OtpSecret: uuid.New().String(),
	}

	err := a.repo.UpdateAccountOTPSecret(ctx, user, opts)
	if err != nil {
		return "", err
	}
	return user.OtpSecret, nil
}
