package auth

import "time"

type Core struct {
	ID        uint
	Name      string
	Email     string `validate:"required,email"`
	Password  string `validate:"required"`
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ServiceInterface interface {
	Login(input Core) (data Core, token string, err error)
}

type RepositoryInterface interface {
	FindUser(email string) (result Core, err error)
}
