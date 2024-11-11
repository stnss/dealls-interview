package user

import (
	"context"
	"github.com/stnss/dealls-interview/internal/entity"
)

type Repository interface {
	GetUserById(ctx context.Context, userID string) (*entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	CreateUser(ctx context.Context, data *entity.User) error
}
