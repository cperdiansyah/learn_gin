package main

import (
	"time"

	serve "github.com/cperdiansyah/gin-rest-learn/api/server"
	routes "github.com/cperdiansyah/gin-rest-learn/api/server/router"
	"github.com/cperdiansyah/gin-rest-learn/configs"
	"github.com/cperdiansyah/gin-rest-learn/internal/app/rest_api/database"
	"github.com/cperdiansyah/gin-rest-learn/internal/app/rest_api/handlers"
	"github.com/cperdiansyah/gin-rest-learn/internal/app/rest_api/repositories"
	"github.com/cperdiansyah/gin-rest-learn/internal/app/rest_api/services"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	config := configs.NewConfig()

	client, err := database.NewSQLClient(database.Config{
		DBDriver:          config.Database.DatabaseDriver,
		DBSource:          config.Database.DatabaseSource,
		MaxOpenConns:      25,
		MaxIdleConns:      25,
		ConnMaxIdleTime:   15 * time.Minute,
		ConnectionTimeout: 5 * time.Second,
	})

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize database client")
		return
	}

	defer func() {
		if err := client.Close(); err != nil {
			log.Error().Msgf("Failed to close database client: %v", err)
		}
	}()

	// Initialize repositories
	userRepo := repositories.NewUserRepository(client.DB)

	//Initialize services
	userService := services.NewUserService(userRepo)

	// Pass services to handlers
	userHandler := handlers.NewUserHandler(userService)

	cors := config.CorsNew()

	router := gin.Default()
	router.Use(cors)

	// Register routes
	routes.RegisterPublicEndpoints(router, userHandler)

	server := serve.NewServer(log.Logger, router, config)
	server.Serve()
}
