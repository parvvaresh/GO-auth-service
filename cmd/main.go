package main

import (
	"auth-service/internal/config"
	"auth-service/internal/handler"
	"auth-service/internal/infrastructure"
	"auth-service/internal/repository"
	"auth-service/internal/router"
	"auth-service/internal/service"
	"auth-service/pkg/sms"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	db := infrastructure.NewPostgres(cfg.DBURL)

	userRepo := repository.NewUserRepository(db)
	profileRepo := repository.NewProfileRepository(db)
	otpRepo := repository.NewOTPRepository(db)

	smsSender := sms.NewMockSender()

	authService := service.NewAuthService(userRepo, profileRepo, otpRepo, smsSender, cfg.JWTSecret)
	profileService := service.NewProfileService(profileRepo)

	authHandler := handler.NewAuthHandler(authService)
	profileHandler := handler.NewProfileHandler(profileService)

	r := gin.Default()

	router.SetupRoutes(r, authHandler, profileHandler, cfg.JWTSecret)

	log.Printf("server started on :%s\n", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
