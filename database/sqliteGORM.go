package database

import (
	"errors"
	"extia/configs"
	appLogger "extia/logger"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	Db                                 *gorm.DB
	errEmptyArgs                       = errors.New("there is no path to sqlite file")
	errCannotPingToDB                  = errors.New("cannot ping database")
	errCannotOpenGORMConnection        = errors.New("cannot open the GORM connection to database")
	errCannotGetUnderlyingDBConnection = errors.New("cannot get the underlying connection to GORM connection")
)

// Instance a DB connection using GORM logging to the the general file
func New() {
	if len(configs.Database.Path) == 0 {
		panic(errEmptyArgs)
	}

	gormLogger := logger.New(appLogger.Logger, logger.Config{
		LogLevel: logger.Error,
		Colorful: false,
	})

	db, err := gorm.Open(
		sqlite.Open(configs.Database.Path),
		&gorm.Config{Logger: gormLogger},
	)
	if err != nil {
		panic(errCannotOpenGORMConnection)
	}

	sqlDb, err := db.DB()
	if err != nil {
		panic(errCannotGetUnderlyingDBConnection)
	}

	if err = sqlDb.Ping(); err != nil {
		panic(errCannotPingToDB)
	}

	Db = db
}
