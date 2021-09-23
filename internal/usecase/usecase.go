package usecase

type UseCase interface {
	UserService
	AuthService
}

type useCase struct {
	UserService
	AuthService
}

func NewUseCase(userService UserService, authService AuthService) UseCase {
	return &useCase{
		UserService: userService,
		AuthService: authService,
	}
}
