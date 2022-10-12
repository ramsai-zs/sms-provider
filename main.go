package main

import (
	"developer.zopsmart.com/go/gofr/cmd/gofr/migration"
	dbmigration "developer.zopsmart.com/go/gofr/cmd/gofr/migration/dbMigration"
	"developer.zopsmart.com/go/gofr/examples/using-postgres/migrations"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

func main() {
	app := gofr.New()

	appName := app.Config.Get("APP_NAME")

	err := migration.Migrate(appName, dbmigration.NewGorm(app.GORM()), migrations.All(), dbmigration.UP, app.Logger)
	if err != nil {
		app.Logger.Error(err)
	}

	app.Start()
}
