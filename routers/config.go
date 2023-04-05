package routers

import (
	"fmt"
	"net/http"
	"ourstartup/config"
	"ourstartup/middlewares/authMiddleware"
	"ourstartup/services/campaign"
	"ourstartup/services/payment"
	"ourstartup/services/transaction"
	"ourstartup/services/user"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type router struct {
	config config.Config
	db     *gorm.DB
}

func Init(config config.Config, db *gorm.DB) *router {
	return &router{config, db}
}

func (r *router) RunRouter() {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{r.config.ClientLink}
	router.Use(cors.New(corsConfig))

	// static routes
	router.StaticFS("/images", http.Dir("images"))

	apiV1 := router.Group("/api/v1")

	// dependencies
	userRepository := user.CreateRepository(r.db)
	campaignRepository := campaign.CreateRepository(r.db)
	transactionRepo := transaction.CreateRepository(r.db)

	userService := user.CreateService(userRepository)
	campaignService := campaign.CreateService(campaignRepository)
	transactionService := transaction.CreateService(transactionRepo)
	paymentService := payment.CreateService(r.config, transactionService, campaignService)
	authService := authMiddleware.CreateService(r.config.JWTSecret, r.config.EncryptionSecret)

	authMiddleware := authMiddleware.CreateAuthMiddleware(authService, userService)

	// init services routers
	// user routes
	userRouters := CreateUserRouter(r, apiV1)
	userRouters.InitRouter(userService, authService, authMiddleware)

	// campaign routes
	campaignRouters := CreateCampaignRouter(r, apiV1)
	campaignRouters.InitRouter(campaignService, userService, authMiddleware)

	// transaction routes
	transactionRouters := CreateTransactionRouters(r, apiV1)
	transactionRouters.InitRouter(transactionService, userService, campaignService, paymentService, authMiddleware)

	err := router.Run(fmt.Sprintf(":%s", r.config.Port))
	if err != nil {
		panic("[Error] failed to start Gin server due to: " + err.Error())
	}
}
