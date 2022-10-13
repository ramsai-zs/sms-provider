package main

import (
	"developer.zopsmart.com/go/gofr/cmd/gofr/migration"
	dbmigration "developer.zopsmart.com/go/gofr/cmd/gofr/migration/dbMigration"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	msghandler "sms-provider/handlers/message"
	provhandler "sms-provider/handlers/provider"
	"sms-provider/migrations"
	"sms-provider/stores/message"
	"sms-provider/stores/provider"
)

func main() {
	app := gofr.New()

	appName := app.Config.Get("APP_NAME")

	err := migration.Migrate(appName, dbmigration.NewGorm(app.GORM()), migrations.All(), dbmigration.UP, app.Logger)
	if err != nil {
		app.Logger.Error(err)
	}

	// initialize stores
	providerStore := provider.New()
	messageStore := message.New()

	// initialize handlers
	providerHandler := provhandler.New(providerStore)
	messageHandler := msghandler.New(messageStore)

	// endpoints for the provider
	app.POST("/provider", providerHandler.Create)
	app.GET("/provider/{id}", providerHandler.GetByID)
	app.PUT("/provider/{id}", providerHandler.Update)
	app.DELETE("/provider/{id}", providerHandler.Delete)

	// endpoints for the message
	app.POST("/message", messageHandler.Create)
	app.GET("/message", messageHandler.Get)
	app.GET("/message/{msgRefID}", messageHandler.GetByID)
	app.PUT("/message/{msgRefID}", messageHandler.Update)
	app.DELETE("/message/{phoneNumber}", messageHandler.Delete)

	app.Start()
}
