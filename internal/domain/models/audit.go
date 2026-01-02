package models

// AuditLog represents an audit trail entry
type AuditLog struct {
	ID          string
	UserID      string
	Action      string
	Resource    string
	ResourceID  string
	Method      string
	Path        string
	IPAddress   string
	UserAgent   string
	StatusCode  int
	BeforeState string // JSON
	AfterState  string // JSON
	Metadata    string // JSON
	CreatedAt   string
}

// AuditAction represents audit action types
type AuditAction string

const (
	ActionCreate AuditAction = "create"
	ActionRead   AuditAction = "read"
	ActionUpdate AuditAction = "update"
	ActionDelete AuditAction = "delete"
	ActionLogin  AuditAction = "login"
	ActionLogout AuditAction = "logout"
)
