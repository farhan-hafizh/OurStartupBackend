package server

import (
	"fmt"
	"log"
	"ourstartup/config"
	"ourstartup/routers"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init(mode string) {
	// load env file
	configData, _ := config.LoadConfig(mode)
	fmt.Println(configData)
	// connect to db
	dsn := configData.DBConnection
	fmt.Println(dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Connected to database!")
	// init routers then run routers
	router := routers.Init(configData, db)
	router.RunRouter()
}
