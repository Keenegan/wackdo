package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCategoryIsValid(t *testing.T) {
	tests := []struct {
		name     string
		category Category
		valid    bool
	}{
		{
			name:     "food is valid",
			category: CategoryFood,
			valid:    true,
		},
		{
			name:     "drink is valid",
			category: CategoryDrink,
			valid:    true,
		},
		{
			name:     "invalid category",
			category: Category("dessert"),
			valid:    false,
		},
		{
			name:     "empty category",
			category: Category(""),
			valid:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.category.IsValid()
			assert.Equal(t, tt.valid, result)
		})
	}
}
