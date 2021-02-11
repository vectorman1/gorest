package db

import (
	"github.com/dystopia-systems/alaskalog"
	"gorest/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

func InitAndMigrate() (*gorm.DB, error) {
	connString := os.Getenv("MYSQL_CONN_STRING")

	alaskalog.Logger.Infoln("Opening MySQL connection...")
	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{
		Logger: logger.New(nil, logger.Config{LogLevel: logger.Silent}),
	})

	if err != nil {
		return nil, err
	}

	alaskalog.Logger.Infoln("MySql connection successful.")

	err = autoMigrate(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func autoMigrate(db *gorm.DB) error {
	alaskalog.Logger.Println("Starting migration...")

	err := db.AutoMigrate(entity.Recipe{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(entity.User{})
	if err != nil {
		return err
	}

	alaskalog.Logger.Println("Migration finished.")

	return nil
}
