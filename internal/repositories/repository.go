package repositories

import (
	"github.com/stnss/dealls-interview/internal/dependencies"
	"github.com/stnss/dealls-interview/internal/repositories/user"
)

type Repository struct {
	UserRepository user.Repository
}

func NewRepository(dep *dependencies.Dependency) *Repository {
	return &Repository{
		UserRepository: user.NewUserRepository(dep.DB),
	}
}
