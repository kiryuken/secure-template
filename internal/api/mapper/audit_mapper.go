package mapper

// Audit log mapper - converts between audit models and DTOs

import (
	"github.com/kiryuken/max-secure/internal/api/dto"
	"github.com/kiryuken/max-secure/internal/domain/models"
)

// AuditLogToDTO converts AuditLog model to AuditLogDTO
func AuditLogToDTO(log *models.AuditLog) dto.AuditLogDTO {
	// TODO: Implement conversion
	return dto.AuditLogDTO{}
}

// AuditLogsToDTOs converts slice of AuditLog models to slice of AuditLogDTOs
func AuditLogsToDTOs(logs []models.AuditLog) []dto.AuditLogDTO {
	// TODO: Implement conversion
	return []dto.AuditLogDTO{}
}
