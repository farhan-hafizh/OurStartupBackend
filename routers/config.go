package routers

import (
	"ourstartup/serverConfig"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Routers interface {
	RunRouter()
}

type router struct {
	config serverConfig.Config
	db     *gorm.DB
}

func Init(config serverConfig.Config, db *gorm.DB) *router {
	return &router{config, db}
}

func (r *router) RunRouter() {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	apiV1 := router.Group("/api/v1")

	userRouters := CreateUserRouter(r.db, apiV1, r.config)
	userRouters.InitRoutes()

	router.Run()
}
