package repositories

import "github.com/stnss/dealls-interview/internal/dependencies"

type Repository struct {
}

func NewRepository(dep *dependencies.Dependency) *Repository {
	return &Repository{}
}
