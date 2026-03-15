package router

import (
	"auth-service/internal/handler"
	"auth-service/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	r *gin.Engine,
	authHandler *handler.AuthHandler,
	profileHandler *handler.ProfileHandler,
	jwtSecret string,
) {
	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/send-otp", authHandler.SendOTP)
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
		}

		profile := api.Group("/profile")
		profile.Use(middleware.Auth(jwtSecret))
		{
			profile.GET("/me", profileHandler.Me)
			profile.PUT("/me", profileHandler.UpdateMe)
		}
	}
}
