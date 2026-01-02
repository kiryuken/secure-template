package mapper

// User mapper - converts between domain models and DTOs

import (
	"github.com/kiryuken/max-secure/internal/api/dto"
	"github.com/kiryuken/max-secure/internal/domain/models"
)

// UserToDTO converts User model to UserDTO
func UserToDTO(user *models.User) dto.UserDTO {
	// TODO: Implement conversion
	return dto.UserDTO{}
}

// UsersToDTOs converts slice of User models to slice of UserDTOs
func UsersToDTOs(users []models.User) []dto.UserDTO {
	// TODO: Implement conversion
	return []dto.UserDTO{}
}

// CreateUserRequestToModel converts CreateUserRequest to User model
func CreateUserRequestToModel(req *dto.CreateUserRequest) *models.User {
	// TODO: Implement conversion
	return &models.User{}
}

// UpdateUserRequestToModel updates User model from UpdateUserRequest
func UpdateUserRequestToModel(user *models.User, req *dto.UpdateUserRequest) {
	// TODO: Implement update
}
