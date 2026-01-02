package mapper

// File mapper - converts between file models and DTOs

import (
	"github.com/kiryuken/max-secure/internal/api/dto"
	"github.com/kiryuken/max-secure/internal/domain/models"
)

// FileToDTO converts File model to FileDTO
func FileToDTO(file *models.File) dto.FileDTO {
	// TODO: Implement conversion
	return dto.FileDTO{}
}

// FilesToDTOs converts slice of File models to slice of FileDTOs
func FilesToDTOs(files []models.File) []dto.FileDTO {
	// TODO: Implement conversion
	return []dto.FileDTO{}
}

// FileToUploadResponse converts File model to FileUploadResponse
func FileToUploadResponse(file *models.File) dto.FileUploadResponse {
	// TODO: Implement conversion
	return dto.FileUploadResponse{}
}
