package database

import (
	"fmt"

	"github.com/OmarAouini/employee-manager/config"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	logrus.Info("connecting db...")
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.CONFIG.DatabaseUser,
		config.CONFIG.DatabasePass,
		config.CONFIG.DatabaseHost,
		config.CONFIG.DatabasePort,
		config.CONFIG.DatabaseName,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Fatal(err)
	}
	dbConfig, _ := db.DB()
	dbConfig.SetMaxIdleConns(config.CONFIG.DatabaseIdleConn)
	dbConfig.SetMaxOpenConns(config.CONFIG.DatabaseMaxConn)

	logrus.Info("connect db OK")

	DB = db
}
