package dbr

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
)

// Init database.

var Db *gorm.DB

func Init(dbURI string, logger *logrus.Logger) *gorm.DB {
	db, err := gorm.Open("mysql", dbURI)

	if err != nil {
		log.Fatalf("coudn't open database: %s", err.Error())
		return nil
	}

	db.SetLogger(&gormLogger{
		name:   "dbLogger",
		logger: logger,
	})
	Db = db.LogMode(true)
	return Db
}

func GetDb() *gorm.DB {
	return Db
}
