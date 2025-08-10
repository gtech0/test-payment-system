package router

import (
	"test-payment-system/controller"
	"test-payment-system/errs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// NewRouter creates a pre-configured gin http router
// with set routes
func NewRouter() *gin.Engine {
	router := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true

	router.Use(cors.New(corsConfig))
	router.Use(errs.ErrorHandler)

	walletController := controller.NewWalletController()
	walletRouter := router.Group("/api/wallet")
	{
		walletRouter.GET("/:address/balance", walletController.GetBalance)
	}

	transactionController := controller.NewTransactionController()
	transactionRouter := router.Group("/api")
	{
		transactionRouter.POST("/send", transactionController.Send)
		transactionRouter.GET("/transactions", transactionController.GetLast)
	}

	// http://localhost:8001/swagger/index.html
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
