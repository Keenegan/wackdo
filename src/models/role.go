package models

type Role string

const (
	RoleManager  Role = "manager"
	RoleEmployee Role = "employee"
	RolePrep     Role = "prep"
	RoleAdmin    Role = "admin"
)

func IsValidRole(role Role) bool {
	switch role {
	case RoleManager, RoleEmployee, RolePrep, RoleAdmin:
		return true
	default:
		return false
	}
}
