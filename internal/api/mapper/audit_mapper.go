package mapper

// Audit log mapper - converts between audit models and DTOs

import (
	"github.com/kiryuken/secure-template/internal/api/dto"
	"github.com/kiryuken/secure-template/internal/domain/models"
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
