package routers

import (
	"ourstartup/handlers"
	"ourstartup/services/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRouters(group *gin.RouterGroup, db *gorm.DB) {
	userRepository := user.CreateRepository(db)
	userService := user.CreateService(userRepository)
	userHandler := handlers.CreateUserHandler(userService)

	user := group.Group("users")
	user.POST("/create", userHandler.RegisterUser)
	user.POST("/login", userHandler.Login)
}
