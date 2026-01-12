package models

type Role string

const (
	RoleManager Role = "manager"
	RoleCashier Role = "cashier"
	RolePrep    Role = "prep"
	RoleAdmin   Role = "admin"
)

func IsValidRole(role Role) bool {
	switch role {
	case RoleManager, RoleCashier, RolePrep, RoleAdmin:
		return true
	default:
		return false
	}
}
