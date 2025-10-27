package routes

import (
	"github.com/cperdiansyah/gin-rest-learn/internal/app/rest_api/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterPublicEndpoints(router *gin.Engine, userHandlers *handlers.User) {
	router.GET("/users", userHandlers.GetAllUsers)
	router.GET("/users/:id", userHandlers.GetUser)
	router.POST("/users", userHandlers.CreateUser)
	router.PUT("/users/:id", userHandlers.UpdateUser)
	router.DELETE("/users/:id", userHandlers.DeleteUser)
}
