package storage

import (
	"log"

	config "todo-list/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// To initialize database connection
func NewDB(params ...string) *gorm.DB {
	var err error
	conString := config.GetMySQLConnectionString()

	DB, err = gorm.Open(mysql.Open(conString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Panic(err)
	}

	return DB
}

// to get the DB instance
func GetDBInstance() *gorm.DB {
	return DB
}
