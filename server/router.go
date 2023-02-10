package server

import (
	"ourstartup/routers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	apiV1 := router.Group("/api/v1")

	routers.UserRouters(apiV1, db)

	return router
}
