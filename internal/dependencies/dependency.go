package dependencies

import (
	"github.com/stnss/dealls-interview/internal/appctx"
	"github.com/stnss/dealls-interview/internal/bootstrap"
	"github.com/stnss/dealls-interview/pkg/databasex"
)

type Dependency struct {
	DB databasex.Adapter
}

func NewDependency(cfg *appctx.Config) *Dependency {
	return &Dependency{
		DB: bootstrap.RegistryDatabase(cfg.DBRead),
	}
}
