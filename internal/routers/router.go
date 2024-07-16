package routers

import (
	"github.com/gin-gonic/gin"
	api2 "github.com/mostafababaii/gorest/internal/handlers/v1/api"
	"github.com/mostafababaii/gorest/internal/middlewares"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	router := gin.New()
	router.HandleMethodNotAllowed = true

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	apiv1 := router.Group("/api/v1")
	authRouters := apiv1.Group("/auth")

	authHandler := api2.NewAuthHandler()
	authRouters.POST("/register", authHandler.Register)
	authRouters.POST("/login", authHandler.Login)

	authMiddleware := middlewares.AuthMiddleware(authHandler.TokenService)

	userRouters := apiv1.Group("/users")
	userHandler := api2.NewUserHandler()
	userRouters.Use(authMiddleware).GET("/", userHandler.Profile)

	return router
}
