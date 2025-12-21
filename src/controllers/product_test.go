package controllers_test

import (
	"testing"

	"wackdo/src/controllers"

	"github.com/stretchr/testify/assert"
)

func TestValidateProductPostRequest(t *testing.T) {
	tests := []struct {
		name    string
		req     controllers.ProductPostRequest
		wantErr bool
		errMsg  string
	}{
		{
			name: "valid request",
			req: controllers.ProductPostRequest{
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
			req: controllers.ProductPostRequest{
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
			req: controllers.ProductPostRequest{
				Name:        "Burger",
				BasePrice:   0,
				Description: "A burger",
				Image:       "link to an image",
				Category:    "drink",
			},
			wantErr: true,
			errMsg:  "base price can't be 0 or less",
		},
		{
			name: "base price < 0",
			req: controllers.ProductPostRequest{
				Name:        "Burger",
				BasePrice:   -12.2,
				Description: "A burger",
				Image:       "link to an image",
				Category:    "drink",
			},
			wantErr: true,
			errMsg:  "base price can't be 0 or less",
		},
		{
			name: "invalid category",
			req: controllers.ProductPostRequest{
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
			err := controllers.ValidateProductPostRequest(&tt.req)

			if tt.wantErr {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.errMsg)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}