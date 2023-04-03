package routers

import (
	"ourstartup/handlers"
	"ourstartup/middlewares/authMiddleware"
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

func (ur *transactionRouters) InitRouter(service transaction.Service, userService user.Service, middleware authMiddleware.Middlerware) {
	handler := handlers.CreateTransactionHandler(service, userService)

	transaction := ur.group.Group("transaction")

	transaction.POST("/create", middleware.GetAuthMiddleware(), handler.CreateTransaction)

}
