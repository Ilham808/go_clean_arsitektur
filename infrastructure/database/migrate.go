package database

import (
	structUser "GoClean/domain/user"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&structUser.User{}); err != nil {
		return err
	}

	//Tambahkan struct lainnya jika ada yang ingin dimigrate

	return nil
}
