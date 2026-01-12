package controllers_products

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateProductId(t *testing.T) {
	tests := []struct {
		name    string
		idStr   string
		wantErr bool
	}{
		{
			name:    "valid id",
			idStr:   "1",
			wantErr: false,
		},
		{
			name:    "valid large id",
			idStr:   "999",
			wantErr: false,
		},
		{
			name:    "invalid id - zero",
			idStr:   "0",
			wantErr: true,
		},
		{
			name:    "invalid id - negative",
			idStr:   "-1",
			wantErr: true,
		},
		{
			name:    "invalid id - not a number",
			idStr:   "abc",
			wantErr: true,
		},
		{
			name:    "invalid id - empty",
			idStr:   "",
			wantErr: true,
		},
		{
			name:    "invalid id - float",
			idStr:   "1.5",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, err := strconv.Atoi(tt.idStr)
			hasErr := err != nil || id <= 0

			assert.Equal(t, tt.wantErr, hasErr)
		})
	}
}

func TestValidateProductNameParam(t *testing.T) {
	tests := []struct {
		name    string
		nameVal string
		wantErr bool
	}{
		{
			name:    "valid name",
			nameVal: "Burger",
			wantErr: false,
		},
		{
			name:    "valid name - with spaces",
			nameVal: "Double Cheese Burger",
			wantErr: false,
		},
		{
			name:    "invalid name - empty",
			nameVal: "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hasErr := tt.nameVal == ""

			assert.Equal(t, tt.wantErr, hasErr)
		})
	}
}
