package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupTestDB() (*gorm.DB, *gorm.DB) {
	dsn := "root:mysql@tcp(127.0.0.1:3306)/go_clean?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	return db, db
}
