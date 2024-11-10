package migration

import (
	"github.com/stnss/dealls-interview/internal/appctx"
	"github.com/stnss/dealls-interview/pkg/databasex"
)

func MigrateDatabase() {
	cfg := appctx.NewConfig()

	databasex.DatabaseMigration(&databasex.Config{
		Driver:       cfg.DBWrite.Driver,
		Host:         cfg.DBWrite.Host,
		Port:         cfg.DBWrite.Port,
		Name:         cfg.DBWrite.Name,
		User:         cfg.DBWrite.User,
		Password:     cfg.DBWrite.Pass,
		Charset:      cfg.DBWrite.Charset,
		DialTimeout:  cfg.DBWrite.DialTimeout,
		MaxIdleConns: cfg.DBWrite.MaxIdle,
		MaxOpenConns: cfg.DBWrite.MaxOpen,
		MaxLifetime:  cfg.DBWrite.MaxLifeTime,
		TimeZone:     cfg.DBWrite.Timezone,
	})
}
