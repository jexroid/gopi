package database

import (
	"fmt"
	"github.com/jexroid/gopi/api"
	"github.com/jexroid/gopi/api/models"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func Init() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_DATABASE"), os.Getenv("DB_PORT"))
	var db, err = gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{
			SkipDefaultTransaction: true,
			PrepareStmt:            true,
		},
	)
	migratingAuthError := db.Table("User").AutoMigrate(&api.User{})
	migratingCarError := db.Table("cars").AutoMigrate(&models.Cars{})
	if migratingAuthError != nil {
		logrus.Panic("migration error : ", migratingAuthError)
	}
	if migratingCarError != nil {
		logrus.Panic("migration error : ", migratingCarError)
	}
	if err != nil {
		logrus.Panic(err)
	}

	return db
}
