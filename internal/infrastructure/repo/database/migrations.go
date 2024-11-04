package database

import (
	"database/sql"
	"os"
	"path"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/labstack/gommon/log"
	"github.com/pressly/goose/v3"
)

func RunMigrations(dsn, migrationsDir string) error {
	curDir, _ := os.Getwd()
	log.Infof("Migrations dir: %s", path.Join(curDir, migrationsDir))

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return err
	}

	defer func() {
		_ = db.Close()
	}()

	if err = db.Ping(); err != nil {
		return err
	}

	if _, err = os.Stat(migrationsDir); err != nil {
		return err
	}

	if err = goose.SetDialect("postgres"); err != nil {
		return err
	}

	return goose.Up(db, migrationsDir)
}
