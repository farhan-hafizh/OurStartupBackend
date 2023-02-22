package server

import (
	"fmt"
	"log"
	"ourstartup/routers"
	"ourstartup/serverConfig"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init() {
	config, _ := serverConfig.LoadConfig()
	dsn := config.DBConnection
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Connected to database!")

	router := routers.Init(config, db)
	router.RunRouter()
}
