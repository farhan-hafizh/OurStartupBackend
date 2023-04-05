package routers

import (
	"ourstartup/handlers"
	"ourstartup/middlewares/authMiddleware"
	"ourstartup/services/campaign"
	"ourstartup/services/transaction"
	"ourstartup/services/user"

	"github.com/gin-gonic/gin"
)

type transactionRouters struct {
	router *router
	group  *gin.RouterGroup
}

func CreateTransactionRouters(router *router, group *gin.RouterGroup) *transactionRouters {
	return &transactionRouters{router, group}
}

func (ur *transactionRouters) InitRouter(service transaction.Service, userService user.Service, campaignService campaign.Service, middleware authMiddleware.Middlerware) {
	handler := handlers.CreateTransactionHandler(service, userService, campaignService)

	transaction := ur.group.Group("transaction")

	transaction.GET("/:slug/:campaignOwner", middleware.GetAuthMiddleware(), handler.GetTransHistoryByCampaign)
	transaction.GET("/:slug", middleware.GetAuthMiddleware(), handler.GetTransHistoryByCampaign)
	transaction.GET("/", middleware.GetAuthMiddleware(), handler.GetTransactionHistory)
	transaction.POST("/create", middleware.GetAuthMiddleware(), handler.CreateTransaction)

}
