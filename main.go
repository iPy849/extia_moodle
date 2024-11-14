package main

import (
	"extia/app"
	"extia/configs"
	"extia/database"
	myLogger "extia/logger"
	"extia/repository"
)

func main() {
	configs.InitializeConf(".env")

	myLogger.New(
		myLogger.STDOUT_LOGGER,
		myLogger.DEFAULT_LOG_FILE,
	)

	database.New()
	if err := database.Db.AutoMigrate(repository.RegisteredModels...); err != nil {
		panic(err)
	}

	app.RunApp()
}
