package cmd

import (
	"context"
	"github.com/spf13/cobra"
	"github.com/stnss/dealls-interview/cmd/http"
	"github.com/stnss/dealls-interview/cmd/migration"
	"github.com/stnss/dealls-interview/pkg/logger"
	"os"
	"os/signal"
	"syscall"
)

func Run() {
	rootCmd := &cobra.Command{}
	ctx, cancel := context.WithCancel(context.Background())
	logger.SetJSONFormatter()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-quit
		cancel()
	}()

	migrateCmd := &cobra.Command{
		Use:   "db:migrate",
		Short: "database migration",
		Run: func(c *cobra.Command, args []string) {
			migration.MigrateDatabase()
		},
	}

	migrateCmd.Flags().BoolP("version", "", false, "print version")
	migrateCmd.Flags().StringP("dir", "", "database/migration/", "directory with migration files")
	migrateCmd.Flags().StringP("table", "", "db", "migrations table name")
	migrateCmd.Flags().BoolP("verbose", "", false, "enable verbose mode")
	migrateCmd.Flags().BoolP("guide", "", false, "print help")
	migrateCmd.Flags().StringP("dsn", "", "", "database dsn")

	httpCmd := &cobra.Command{
		Use:   "http",
		Short: "Start HTTP server",
		Run: func(cmd *cobra.Command, args []string) {
			http.Start(ctx)
		},
	}

	rootCmd.AddCommand(httpCmd, migrateCmd)
	if err := rootCmd.Execute(); err != nil {
		logger.Fatal(err)
	}
}
