package entity

import (
	"context"

	internalerror "github.com/andremelinski/auction/internal/internal_error"
)

type User struct {
	Id   string
	Name string
}

type IUserRepository interface {
	FindUserById(ctx context.Context, id string) (*User, *internalerror.InternalError)
}
