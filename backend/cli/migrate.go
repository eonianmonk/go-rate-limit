package cli

import (
	"fmt"
	"log"

	"github.com/eonianmonk/go-rate-limit/assets"
	"github.com/eonianmonk/go-rate-limit/backend/config"
	migrate "github.com/rubenv/sql-migrate"

	_ "github.com/lib/pq"
)

func MigrateUp(cfg config.Config) error {
	migrations := migrate.EmbedFileSystemMigrationSource{
		FileSystem: assets.Migrations,
		Root:       "migrations",
	}
	applied, err := migrate.Exec(cfg.DB(), cfg.DBType(), migrations, migrate.Up)
	if err != nil {
		panic(fmt.Sprintf("failed to migrate db: %s", err.Error()))
	}
	log.Printf("applied %d migrations\n", applied)
	return nil
}
