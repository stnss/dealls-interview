package dependencies

import (
	"github.com/stnss/dealls-interview/internal/appctx"
	"github.com/stnss/dealls-interview/internal/bootstrap"
	"github.com/stnss/dealls-interview/pkg/cryptx"
	"github.com/stnss/dealls-interview/pkg/databasex"
	"github.com/stnss/dealls-interview/pkg/jwtx"
)

type Dependency struct {
	DB     databasex.Adapter
	Jwt    jwtx.Helper
	Crypto cryptx.Helper
}

func NewDependency(cfg *appctx.Config) *Dependency {
	return &Dependency{
		DB:     bootstrap.RegistryDatabase(cfg.DBRead),
		Jwt:    jwtx.NewJwtHelper(),
		Crypto: cryptx.NewCryptox(),
	}
}
