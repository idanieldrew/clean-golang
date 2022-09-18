package cmd

import (
	mysql2 "clean-golang/app/infrastructure/database/mysql"
	"clean-golang/app/infrastructure/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/urfave/cli/v2"
)

func Migrate(cCtx *cli.Context) error {
	driver, err := mysql.WithInstance(mysql2.Db, &mysql.Config{})
	if err != nil {
		logger.Error("problem in migration")
		return err
	}

	m, err := migrate.NewWithDatabaseInstance("file://migrations", "mysql", driver)
	if err != nil {
		logger.Error("problem in migration 2")
		return err
	}

	_ = m.Steps(2)
	return nil
}
