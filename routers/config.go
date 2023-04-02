package routers

import (
	"fmt"
	"net/http"
	"ourstartup/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Routers interface {
	RunRouter()
}

type router struct {
	config config.Config
	db     *gorm.DB
}

func Init(config config.Config, db *gorm.DB) *router {
	return &router{config, db}
}

func (r *router) RunRouter() {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// static routes
	router.StaticFS("/images", http.Dir("images"))

	apiV1 := router.Group("/api/v1")

	// user routes
	userRouters := CreateUserRouter(r, apiV1)
	userRouters.InitRouter()

	// campaign routes
	campaignRouters := CreateCampaignRouter(r, apiV1)
	campaignRouters.InitRouter()

	err := router.Run(fmt.Sprintf(":%s", r.config.Port))
	if err != nil {
		panic("[Error] failed to start Gin server due to: " + err.Error())
	}
}
