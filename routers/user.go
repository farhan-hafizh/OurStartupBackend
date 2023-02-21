package routers

import (
	"ourstartup/handlers"
	"ourstartup/middlewares/auth"
	"ourstartup/services/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRouters(group *gin.RouterGroup, db *gorm.DB) {
	userRepository := user.CreateRepository(db)
	userService := user.CreateService(userRepository)
	authService := auth.CreateService()
	userHandler := handlers.CreateUserHandler(userService, authService)

	user := group.Group("users")
	user.POST("/create", userHandler.RegisterUser)
	user.POST("/login", userHandler.Login)
	user.POST("/check-email", userHandler.CheckEmailAvailability)
	user.POST("/upload-avatar", userHandler.UploadAvatar)
}
