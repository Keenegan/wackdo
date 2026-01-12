package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidRole(t *testing.T) {
	tests := []struct {
		name  string
		role  Role
		valid bool
	}{
		{
			name:  "manager is valid",
			role:  RoleManager,
			valid: true,
		},
		{
			name:  "employee is valid",
			role:  RoleEmployee,
			valid: true,
		},
		{
			name:  "prep is valid",
			role:  RolePrep,
			valid: true,
		},
		{
			name:  "admin is valid",
			role:  RoleAdmin,
			valid: true,
		},
		{
			name:  "invalid role",
			role:  Role("superadmin"),
			valid: false,
		},
		{
			name:  "empty role",
			role:  Role(""),
			valid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidRole(tt.role)
			assert.Equal(t, tt.valid, result)
		})
	}
}
