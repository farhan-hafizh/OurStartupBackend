package routers

import (
	"ourstartup/handlers"
	"ourstartup/middlewares/authMiddleware"
	"ourstartup/services/campaign"
	"ourstartup/services/user"

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

	userRepository := user.CreateRepository(ur.router.db)
	userService := user.CreateService(userRepository)

	authService := authMiddleware.CreateService(ur.router.config.JWTSecret, ur.router.config.EncryptionSecret)
	authMiddleware := authMiddleware.CreateAuthMiddleware(authService, userService)

	handler := handlers.CreateCampaignHandler(service, userService)

	campaign := ur.group.Group("campaign")

	campaign.GET("/", authMiddleware.GetAuthMiddleware(), handler.GetCampaigns)
	campaign.POST("/create", authMiddleware.GetAuthMiddleware(), handler.CreateCampaign)
	campaign.PUT("/:slug", authMiddleware.GetAuthMiddleware(), handler.UpdateCampaign)
	campaign.GET("/:slug", authMiddleware.GetAuthMiddleware(), handler.GetCampaignDetail)
}
