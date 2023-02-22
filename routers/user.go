package routers

import (
	"ourstartup/handlers"
	"ourstartup/middlewares/auth"
	"ourstartup/services/user"

	"github.com/gin-gonic/gin"
)

type UserRouters interface {
	InitRoutes()
}

type userRouters struct {
	router *router
	group  *gin.RouterGroup
}

func CreateUserRouter(router *router, group *gin.RouterGroup) *userRouters {
	return &userRouters{router, group}
}

func (ur *userRouters) InitRouter() {
	userRepository := user.CreateRepository(ur.router.db)
	userService := user.CreateService(userRepository)
	authService := auth.CreateService(ur.router.config.JWTSecret)
	userHandler := handlers.CreateUserHandler(userService, authService)

	user := ur.group.Group("users")
	user.POST("/create", userHandler.RegisterUser)
	user.POST("/login", userHandler.Login)
	user.POST("/check-email", userHandler.CheckEmailAvailability)
	user.POST("/upload-avatar", userHandler.UploadAvatar)
}
