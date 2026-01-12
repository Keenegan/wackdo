package controllers_order

import (
	"testing"

	"wackdo/src/models"

	"github.com/stretchr/testify/assert"
)

func TestOrderStatusIsValid(t *testing.T) {
	tests := []struct {
		name   string
		status models.OrderStatus
		valid  bool
	}{
		{
			name:   "pending is valid",
			status: models.OrderStatusPending,
			valid:  true,
		},
		{
			name:   "confirmed is valid",
			status: models.OrderStatusConfirmed,
			valid:  true,
		},
		{
			name:   "preparing is valid",
			status: models.OrderStatusPreparing,
			valid:  true,
		},
		{
			name:   "ready is valid",
			status: models.OrderStatusReady,
			valid:  true,
		},
		{
			name:   "delivered is valid",
			status: models.OrderStatusDelivered,
			valid:  true,
		},
		{
			name:   "cancelled is valid",
			status: models.OrderStatusCancelled,
			valid:  true,
		},
		{
			name:   "invalid status",
			status: models.OrderStatus("invalid"),
			valid:  false,
		},
		{
			name:   "empty status",
			status: models.OrderStatus(""),
			valid:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.status.IsValid()
			assert.Equal(t, tt.valid, result)
		})
	}
}
