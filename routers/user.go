package routers

import (
	"ourstartup/handlers"
	"ourstartup/middlewares/authMiddleware"
	"ourstartup/services/user"

	"github.com/gin-gonic/gin"
)

type userRouters struct {
	router *router
	group  *gin.RouterGroup
}

func CreateUserRouter(router *router, group *gin.RouterGroup) *userRouters {
	return &userRouters{router, group}
}

func (ur *userRouters) InitRouter(service user.Service, authService authMiddleware.Service, authMiddleware authMiddleware.Middlerware) {
	userHandler := handlers.CreateUserHandler(service, authService)

	user := ur.group.Group("users")
	user.POST("/create", userHandler.RegisterUser)
	user.POST("/login", userHandler.Login)
	user.GET("/fetch", userHandler.FetchUser)
	user.POST("/check-email", userHandler.CheckEmailAvailability)
	user.POST("/upload-avatar", authMiddleware.GetAuthMiddleware(), userHandler.UploadAvatar)
}
