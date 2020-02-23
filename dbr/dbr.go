package dbr

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
)

// Init database.
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
	return db.LogMode(true)
}
