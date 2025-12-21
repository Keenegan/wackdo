package controllers_test

import (
	"testing"

	"wackdo/src/controllers"
	"wackdo/src/models"

	"github.com/stretchr/testify/assert"
)

func TestValidateEmployeeRequest(t *testing.T) {
	tests := []struct {
		name    string
		req     controllers.EmployeePostRequest
		wantErr bool
		errMsg  string
	}{
		{
			name: "valid request",
			req: controllers.EmployeePostRequest{
				Name:  "Alice",
				Roles: []models.Role{models.RoleAdmin},
			},
			wantErr: false,
		},
		{
			name: "missing name",
			req: controllers.EmployeePostRequest{
				Name:  "",
				Roles: []models.Role{models.RoleAdmin},
			},
			wantErr: true,
			errMsg:  "name is required",
		},
		{
			name: "name only spaces",
			req: controllers.EmployeePostRequest{
				Name:  "   ",
				Roles: []models.Role{models.RoleAdmin},
			},
			wantErr: true,
			errMsg:  "name is required",
		},
		{
			name: "missing roles",
			req: controllers.EmployeePostRequest{
				Name:  "Bob",
				Roles: []models.Role{},
			},
			wantErr: true,
			errMsg:  "roles is required",
		},
		{
			name: "invalid role",
			req: controllers.EmployeePostRequest{
				Name:  "Charlie",
				Roles: []models.Role{"boss"},
			},
			wantErr: true,
			errMsg:  "invalid role",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := controllers.ValidateEmployeePostRequest(tt.req)

			if tt.wantErr {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.errMsg)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
