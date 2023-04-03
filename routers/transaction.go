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

func (ur *transactionRouters) InitRouter() {
	transactionRepo := transaction.CreateRepository(ur.router.db)
	transactionService := transaction.CreateService(transactionRepo)

	userRepository := user.CreateRepository(ur.router.db)
	userService := user.CreateService(userRepository)

	authService := authMiddleware.CreateService(ur.router.config.JWTSecret, ur.router.config.EncryptionSecret)
	authMiddleware := authMiddleware.CreateAuthMiddleware(authService, userService)

	handler := handlers.CreateTransactionHandler(transactionService, userService)

	transaction := ur.group.Group("transaction")

	transaction.POST("/create", authMiddleware.GetAuthMiddleware(), handler.CreateTransaction)

}
