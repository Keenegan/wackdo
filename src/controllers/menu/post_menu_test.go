package controllers_menus

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateMenuPostRequest(t *testing.T) {
	tests := []struct {
		name    string
		req     MenuPostRequest
		wantErr bool
		errMsg  string
	}{
		{
			name: "valid request",
			req: MenuPostRequest{
				Name:        "Menu Burger",
				BasePrice:   15.0,
				Description: "Un menu burger",
				Image:       "image.jpg",
				ProductIds:  []uint{1, 2},
			},
			wantErr: false,
		},
		{
			name: "name only spaces",
			req: MenuPostRequest{
				Name:        "   ",
				BasePrice:   15.0,
				Description: "Un menu",
				ProductIds:  []uint{1},
			},
			wantErr: true,
			errMsg:  "name is required",
		},
		{
			name: "empty name",
			req: MenuPostRequest{
				Name:        "",
				BasePrice:   15.0,
				Description: "Un menu",
				ProductIds:  []uint{1},
			},
			wantErr: true,
			errMsg:  "name is required",
		},
		{
			name: "base price 0",
			req: MenuPostRequest{
				Name:        "Menu Test",
				BasePrice:   0,
				Description: "Un menu",
				ProductIds:  []uint{1},
			},
			wantErr: true,
			errMsg:  "basePrice must be greater than 0",
		},
		{
			name: "base price negative",
			req: MenuPostRequest{
				Name:        "Menu Test",
				BasePrice:   -5.0,
				Description: "Un menu",
				ProductIds:  []uint{1},
			},
			wantErr: true,
			errMsg:  "basePrice must be greater than 0",
		},
		{
			name: "empty productIds",
			req: MenuPostRequest{
				Name:        "Menu Test",
				BasePrice:   15.0,
				Description: "Un menu",
				ProductIds:  []uint{},
			},
			wantErr: true,
			errMsg:  "productIds must contain at least one product",
		},
		{
			name: "nil productIds",
			req: MenuPostRequest{
				Name:        "Menu Test",
				BasePrice:   15.0,
				Description: "Un menu",
				ProductIds:  nil,
			},
			wantErr: true,
			errMsg:  "productIds must contain at least one product",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateMenuPostRequest(&tt.req)

			if tt.wantErr {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.errMsg)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
