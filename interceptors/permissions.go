package interceptors

import "slices"

type Permission string

const (
	PermCreateBot Permission = "bot:create"
	PermRunBot    Permission = "bot:run"
	PermViewBot   Permission = "bot:view"
)

var rolePermissions = map[string][]Permission{
	"admin": {PermCreateBot, PermRunBot, PermViewBot},
	"dev":   {PermRunBot, PermViewBot},
	"user":  {PermViewBot},
}

func HasPermission(role string, perm Permission) bool {
	perms, exists := rolePermissions[role]
	if !exists {
		return false
	}
	return slices.Contains(perms, perm)
}
