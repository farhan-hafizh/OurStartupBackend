package routers

import (
	"ourstartup/handlers"
	"ourstartup/middlewares/authMiddleware"
	"ourstartup/services/campaign"

	"github.com/gin-gonic/gin"
)

type CampaignRouters interface {
	InitRoutes()
}

type campaignRouters struct {
	router *router
	group  *gin.RouterGroup
}

func CreateCampaignRouter(router *router, group *gin.RouterGroup) *campaignRouters {
	return &campaignRouters{router, group}
}

func (ur *campaignRouters) InitRouter() {
	repository := campaign.CreateRepository(ur.router.db)
	service := campaign.CreateService(repository)

	authService := authMiddleware.CreateService(ur.router.config.JWTSecret, ur.router.config.EncryptionSecret)
	authMiddleware := authMiddleware.CreateAuthMiddleware(authService)

	handler := handlers.CreateCampaignHandler(service)

	campaign := ur.group.Group("campaign")

	campaign.GET("/", authMiddleware.GetAuthMiddleware(), handler.GetCampaigns)
	campaign.POST("/create", authMiddleware.GetAuthMiddleware(), handler.CreateCampaign)
}
