package routers

import (
	"ourstartup/handlers"
	"ourstartup/middlewares/auth"
	"ourstartup/serverConfig"
	"ourstartup/services/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserRouters interface {
	InitRoutes()
}

type userRouters struct {
	db     *gorm.DB
	group  *gin.RouterGroup
	config serverConfig.Config
}

func CreateUserRouter(db *gorm.DB, group *gin.RouterGroup, config serverConfig.Config) *userRouters {
	return &userRouters{db, group, config}
}

func (r *userRouters) InitRoutes() {
	userRepository := user.CreateRepository(r.db)
	userService := user.CreateService(userRepository)
	authService := auth.CreateService(r.config.JWTSecret)
	userHandler := handlers.CreateUserHandler(userService, authService)

	user := r.group.Group("users")
	user.POST("/create", userHandler.RegisterUser)
	user.POST("/login", userHandler.Login)
	user.POST("/check-email", userHandler.CheckEmailAvailability)
	user.POST("/upload-avatar", userHandler.UploadAvatar)
}
