package controllers_user

import (
	"net/mail"
	"testing"

	"wackdo/src/models"

	"github.com/stretchr/testify/assert"
)

func TestEmailValidation(t *testing.T) {
	tests := []struct {
		name    string
		email   string
		wantErr bool
	}{
		{
			name:    "valid email",
			email:   "test@example.com",
			wantErr: false,
		},
		{
			name:    "valid email with subdomain",
			email:   "user@mail.example.com",
			wantErr: false,
		},
		{
			name:    "invalid email - no @",
			email:   "testexample.com",
			wantErr: true,
		},
		{
			name:    "invalid email - no domain",
			email:   "test@",
			wantErr: true,
		},
		{
			name:    "invalid email - empty",
			email:   "",
			wantErr: true,
		},
		{
			name:    "invalid email - spaces",
			email:   "test @example.com",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := mail.ParseAddress(tt.email)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestRoleValidation(t *testing.T) {
	tests := []struct {
		name  string
		role  models.Role
		valid bool
	}{
		{
			name:  "manager is valid",
			role:  models.RoleManager,
			valid: true,
		},
		{
			name:  "cashier is valid",
			role:  models.RoleCashier,
			valid: true,
		},
		{
			name:  "prep is valid",
			role:  models.RolePrep,
			valid: true,
		},
		{
			name:  "admin is valid",
			role:  models.RoleAdmin,
			valid: true,
		},
		{
			name:  "invalid role",
			role:  models.Role("superadmin"),
			valid: false,
		},
		{
			name:  "empty role",
			role:  models.Role(""),
			valid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := models.IsValidRole(tt.role)
			assert.Equal(t, tt.valid, result)
		})
	}
}
