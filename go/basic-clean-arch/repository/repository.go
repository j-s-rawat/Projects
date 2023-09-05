// repository/repository.go
package repository

import "clean_arch/domain"

type UserRepositoryImpl struct {
	// Simulate a database connection or any other data source
	users []domain.User
}

func NewUserRepositoryImpl() *UserRepositoryImpl {
	return &UserRepositoryImpl{
		users: []domain.User{
			{ID: 1, Username: "user1", Email: "user1@example.com"},
			{ID: 2, Username: "user2", Email: "user2@example.com"},
		},
	}
}

func (r *UserRepositoryImpl) GetByID(id int) (*domain.User, error) {
	for _, user := range r.users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, &domain.ErrUserNotFound{}
}
