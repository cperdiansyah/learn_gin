package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/cperdiansyah/gin-rest-learn/internal/app/rest_api/models/dtos"
	"github.com/cperdiansyah/gin-rest-learn/internal/app/rest_api/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type User struct {
	userService *services.User
}

func NewUserHandler(userService *services.User) *User {
	return &User{userService: userService}
}

func (h *User) GetAllUsers(ctx *gin.Context) {
	allUsers, err := h.userService.GetAllUsers()
	if err != nil {
		ctx.AbortWithStatusJSON(err.Code, err)

		return
	}

	ctx.JSON(http.StatusOK, allUsers)
}

func (h *User) GetUser(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "User ID not valid"})

		return
	}

	user, userErr := h.userService.GetUser(userID)
	if userErr != nil {
		ctx.AbortWithStatusJSON(userErr.Code, userErr)

		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (h *User) DeleteUser(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "User ID not valid"})

		return
	}

	deleteError := h.userService.DeleteUser(userID)
	if deleteError != nil {
		ctx.AbortWithStatusJSON(deleteError.Code, deleteError)

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}

func (h *User) CreateUser(ctx *gin.Context) {
	var createUserRequest dtos.CreateUserRequest

	if err := ctx.ShouldBindJSON(&createUserRequest); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make(map[string]string)
			for _, fe := range ve {
				out[fe.Field()] = msgForTag(fe)
			}
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})

			return
		}
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	createUserResponse, signupError := h.userService.CreateUser(&createUserRequest)
	if signupError != nil {
		ctx.AbortWithStatusJSON(signupError.Code, signupError)

		return
	}

	ctx.JSON(http.StatusCreated, createUserResponse)
}

func (h *User) UpdateUser(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "User ID not valid"})

		return
	}

	var updateUserRequest dtos.UpdateUserRequest

	if err := ctx.ShouldBindJSON(&updateUserRequest); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make(map[string]string)
			for _, fe := range ve {
				out[fe.Field()] = msgForTag(fe)
			}
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})

			return
		}
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	updateError := h.userService.UpdateUser(userID, &updateUserRequest)
	if updateError != nil {
		ctx.AbortWithStatusJSON(updateError.Code, updateError)

		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{"message": "User updated"})
}

func msgForTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "min":
		return fmt.Sprintf("Minimum length is %s", fe.Param())
	case "custom_password":
		return "Password must be at least 8 characters long and include uppercase, lowercase, number, and special character"
	default:
		return "Invalid value"
	}
}
