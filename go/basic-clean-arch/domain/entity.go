// domain/entity.go
package domain

type User struct {
	ID       int
	Username string
	Email    string
}

type ErrUserNotFound struct{}

func (e *ErrUserNotFound) Error() string {
	return "User not found error"
}
