package userusecase

import (
	"context"

	"github.com/andremelinski/auction/internal/entity"
	internalerror "github.com/andremelinski/auction/internal/internal_error"
)

type IUserUserCase interface {
	FindUserById(ctx context.Context, id string) (*UserOutputDTO, *internalerror.InternalError)
}

type UserOutputDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UserUserCase struct {
	UserRepository entity.IUserRepository
}

func NewUserUserCase(userRepository entity.IUserRepository) *UserUserCase {
	return &UserUserCase{
		userRepository,
	}
}

func (uu *UserUserCase) FindUserById(ctx context.Context, id string) (*UserOutputDTO, *internalerror.InternalError) {
	user, err := uu.UserRepository.FindUserById(ctx, id)

	if err != nil {
		return nil, err
	}
	return &UserOutputDTO{
		ID:   user.Id,
		Name: user.Name,
	}, nil
}
