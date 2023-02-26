package server

import (
	"fmt"
	"log"
	"ourstartup/routers"
	"ourstartup/serverConfig"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init(mode string) {
	// load env file
	config, _ := serverConfig.LoadConfig(mode)
	// connect to db
	dsn := config.DBConnection
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Connected to database!")
	// init routers then run routers
	router := routers.Init(config, db)
	router.RunRouter()
}
