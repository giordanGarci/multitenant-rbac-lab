package structs

type Role string

const (
	RoleAdmin  Role = "ADMIN"
	RoleDev    Role = "DEV"
	RoleViewer Role = "VIEWER"
)

type UserContext struct {
	UserID   int64
	TenantID int64
	Role
}
