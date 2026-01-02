package dto

// Audit log DTOs

// AuditLogDTO represents audit log data transfer object
type AuditLogDTO struct {
	ID          string                 `json:"id"`
	UserID      string                 `json:"user_id,omitempty"`
	Username    string                 `json:"username,omitempty"`
	Action      string                 `json:"action"`
	Resource    string                 `json:"resource"`
	ResourceID  string                 `json:"resource_id,omitempty"`
	Method      string                 `json:"method"`
	Path        string                 `json:"path"`
	IPAddress   string                 `json:"ip_address"`
	UserAgent   string                 `json:"user_agent"`
	StatusCode  int                    `json:"status_code"`
	BeforeState map[string]interface{} `json:"before_state,omitempty"`
	AfterState  map[string]interface{} `json:"after_state,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	CreatedAt   string                 `json:"created_at"`
}

// AuditLogListRequest represents audit log list filters
type AuditLogListRequest struct {
	UserID    string `json:"user_id,omitempty"`
	Action    string `json:"action,omitempty"`
	Resource  string `json:"resource,omitempty"`
	StartDate string `json:"start_date,omitempty"`
	EndDate   string `json:"end_date,omitempty"`
	Page      int    `json:"page"`
	PageSize  int    `json:"page_size"`
}

// AuditLogListResponse represents audit log list with pagination
type AuditLogListResponse struct {
	Logs       []AuditLogDTO `json:"logs"`
	Total      int64         `json:"total"`
	Page       int           `json:"page"`
	PageSize   int           `json:"page_size"`
	TotalPages int           `json:"total_pages"`
}
