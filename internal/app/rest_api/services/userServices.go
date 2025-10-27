package services

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/cperdiansyah/gin-rest-learn/internal/app/rest_api/models"
	"github.com/cperdiansyah/gin-rest-learn/internal/app/rest_api/models/dtos"
	"github.com/cperdiansyah/gin-rest-learn/internal/app/rest_api/repositories"
)

type User struct {
	userRepo *repositories.User
}

func NewUserService(userRepo *repositories.User) *User {
	return &User{userRepo: userRepo}
}

func (us *User) GetAllUsers() (*dtos.GetAllUsersResponse, *models.ErrorResponse) {
	response := &dtos.GetAllUsersResponse{}

	queriedUsers, err := us.userRepo.GetAllUsers()
	if err != nil {
		return nil, &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
		}
	}

	response.MapUsersResponse(queriedUsers)

	return response, nil
}

func (us *User) GetUser(userID int) (*dtos.UserResponse, *models.ErrorResponse) {
	response := &dtos.UserResponse{}

	user, err := us.userRepo.FindById(userID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &models.ErrorResponse{
				Code:    http.StatusNotFound,
				Message: "User Not Found",
			}
		}
		return nil, &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
		}
	}

	response.MapUserResponse(user)

	return response, nil
}

func (us *User) DeleteUser(userId int) *models.ErrorResponse {
	user, err := us.userRepo.FindById(userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &models.ErrorResponse{
				Code:    http.StatusNotFound,
				Message: "User not found",
			}
		}
		return &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
		}
	}
	err = us.userRepo.DeleteUser(user.ID)
	if err != nil {
		return &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
		}
	}

	return nil
}

func (us *User) CreateUser(createUserRequest *dtos.CreateUserRequest) (*dtos.CreateUserResponse, *models.ErrorResponse) {
	userResponse := &dtos.CreateUserResponse{}

	errEmail := us.checkIfEmailExists(createUserRequest.Email)
	if errEmail != nil {
		return nil, errEmail
	}

	user := createUserRequest.ToUser()

	err := us.userRepo.Create(user)
	if err != nil {
		return nil, &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to create user",
		}
	}

	return userResponse.FromUser(user), nil
}

func (us *User) UpdateUser(userID int, updateUserRequest *dtos.UpdateUserRequest) *models.ErrorResponse {
	existingUser, err := us.userRepo.FindById(userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &models.ErrorResponse{
				Code:    http.StatusNotFound,
				Message: "User not found",
			}
		}
		return &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
		}
	}

	if updateUserRequest.Email != existingUser.Email {
		errEmail := us.checkIfEmailExists(updateUserRequest.Email)
		if errEmail != nil {
			return errEmail
		}
	}

	existingUser = updateUserRequest.ToUser()
	existingUser.ID = userID

	err = us.userRepo.Update(existingUser)

	if err != nil {
		return &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to update user",
		}
	}

	return nil
}

func (us *User) checkIfEmailExists(email string) *models.ErrorResponse {
	userWithEmail, err := us.userRepo.FindByEmail(email)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
		}
	}
	if userWithEmail != nil {
		return &models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Email already in use",
		}
	}
	return nil
}
