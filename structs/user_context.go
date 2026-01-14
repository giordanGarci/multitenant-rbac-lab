package structs

type Role string

const (
	RoleAdmin  Role = "ADMIN"
	RoleDev    Role = "DEV"
	RoleViewer Role = "VIEWER"
)

type UserContext struct {
	UserID   string
	TenantID string
	Role
}
