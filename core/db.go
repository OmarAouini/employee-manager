package core

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//global db connection
var DB *gorm.DB

//connect to database and set global DB connection
func ConnectDb(username string, password string, host string, port string, dbname string, minConn int, maxConn int) {
	fmt.Printf("connecting to db schema %s...\n", dbname)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbname)
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	dbConf, _ := db.DB()
	dbConf.SetMaxIdleConns(minConn)
	dbConf.SetMaxOpenConns(maxConn)
	DB = db
}

//migrate model to update database tables
func Migrate() {
	DB.AutoMigrate(
		&Company{},
		&Project{},
		&Employee{},
	)
}
