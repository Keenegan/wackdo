package controllers_products

import (
	"testing"


	"github.com/stretchr/testify/assert"
)

func TestValidateProductPostRequest(t *testing.T) {
	tests := []struct {
		name    string
		req     ProductPostRequest
		wantErr bool
		errMsg  string
	}{
		{
			name: "valid request",
			req: ProductPostRequest{
				Name:        "Burger",
				BasePrice:   12.0,
				Description: "A burger",
				Image:       "link to an image",
				Category:    "drink",
			},
			wantErr: false,
		},
		{
			name: "name only spaces",
			req: ProductPostRequest{
				Name:        " ",
				BasePrice:   0,
				Description: "A burger",
				Image:       "link to an image",
				Category:    "drink",
			},
			wantErr: true,
			errMsg:  "name is required",
		},
		{
			name: "base price 0",
			req: ProductPostRequest{
				Name:        "Burger",
				BasePrice:   0,
				Description: "A burger",
				Image:       "link to an image",
				Category:    "drink",
			},
			wantErr: true,
			errMsg:  "basePrice must be greater than 0",
		},
		{
			name: "base price < 0",
			req: ProductPostRequest{
				Name:        "Burger",
				BasePrice:   -12.2,
				Description: "A burger",
				Image:       "link to an image",
				Category:    "drink",
			},
			wantErr: true,
			errMsg:  "basePrice must be greater than 0",
		},
		{
			name: "invalid category",
			req: ProductPostRequest{
				Name:        "Burger",
				BasePrice:   12.2,
				Description: "A burger",
				Image:       "link to an image",
				Category:    "invalid cat",
			},
			wantErr: true,
			errMsg:  "invalid category",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateProductPostRequest(&tt.req)

			if tt.wantErr {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.errMsg)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}