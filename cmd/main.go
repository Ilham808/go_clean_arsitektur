package main

import (
	"GoClean/config"
	"GoClean/domain/user"
	"GoClean/infrastructure/database"
	"GoClean/infrastructure/http"
	userHandler "GoClean/infrastructure/http/user"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db := InitDB(*config.InitConfig())

	// Disini Inisialisasi repository
	userRepository := database.NewUserRepository(db)

	// Disini Inisialisasi usecase, handler, dan router
	userUsecase := user.NewUserUsecase(userRepository)
	userHandler := userHandler.NewUserHandler(userUsecase)

	e := echo.New()
	http.SetRoutesUser(e, userHandler)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", 8000)).Error())
}

func InitDB(config config.ProgramConfig) *gorm.DB {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		fmt.Println(connectionString)
		logrus.Error("Cannot connect to database, ", err.Error())
		return nil

	}
	if err := database.AutoMigrate(db); err != nil {
		panic("failed to auto-migrate database")
	}

	return db
}
