package usecase

type UseCase interface {
	UserService
}

type useCase struct {
	UserService
}

func NewUseCase(userService userService) UseCase {
	return &useCase{
		UserService: userService,
	}
}
