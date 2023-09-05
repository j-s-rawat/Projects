// usecase/usecase.go
package usecase

import "clean_arch/domain"

//go:generate mockgen -source=usecase.go -destination=../mocks/mock_usecase.go
type UserRepository interface {
	GetByID(id int) (*domain.User, error)
}

type UserUseCase struct {
	userRepository UserRepository
}

func NewUserUseCase(userRepository UserRepository) *UserUseCase {
	return &UserUseCase{
		userRepository: userRepository,
	}
}

func (uc *UserUseCase) GetUserByID(id int) (*domain.User, error) {
	return uc.userRepository.GetByID(id)
}
