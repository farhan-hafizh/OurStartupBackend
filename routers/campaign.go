package routers

import (
	"ourstartup/handlers"
	"ourstartup/middlewares/authMiddleware"
	"ourstartup/services/campaign"
	"ourstartup/services/user"

	"github.com/gin-gonic/gin"
)

type campaignRouters struct {
	router *router
	group  *gin.RouterGroup
}

func CreateCampaignRouter(router *router, group *gin.RouterGroup) *campaignRouters {
	return &campaignRouters{router, group}
}

func (ur *campaignRouters) InitRouter(service campaign.Service, userService user.Service, middleware authMiddleware.Middlerware) {
	handler := handlers.CreateCampaignHandler(service, userService)

	campaign := ur.group.Group("campaign")

	campaign.GET("/", middleware.GetAuthMiddleware(), handler.GetCampaigns)
	campaign.POST("/create", middleware.GetAuthMiddleware(), handler.CreateCampaign)
	campaign.POST("/upload-image", middleware.GetAuthMiddleware(), handler.UploadCampaignImage)
	campaign.PUT("/:slug", middleware.GetAuthMiddleware(), handler.UpdateCampaign)
	campaign.GET("/:slug", middleware.GetAuthMiddleware(), handler.GetCampaignDetail)
}
