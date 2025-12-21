package models

type Role string

const (
	RoleAdmin    Role = "admin"
	RoleEmployee Role = "employee"
	RoleManager  Role = "manager"
	RolePrep     Role = "prep"
)

func (r Role) IsValid() bool {
	switch r {
	case RoleAdmin, RoleManager, RoleEmployee, RolePrep:
		return true
	default:
		return false
	}
}

func AreAllRolesValid(r []Role) bool {
	for _, role := range r {
		if !role.IsValid() {
			return false
		}
	}
	return true
}
