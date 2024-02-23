package Usecase

import (
	"testing"

	"github.com/MikeMwita/savannah-ordermanagement/internal/core/models"
)

func TestSendSmsAlert(t *testing.T) {
	// Create a sample order
	order := models.Order{
		CustomerID: 1,
		Item:       "Test Item",
	}

	// Define test cases
	tests := []struct {
		name    string
		order   models.Order
		wantErr bool
	}{
		{
			name:    "Valid order",
			order:   order,
			wantErr: false,
		},
		// Add more test cases if needed
	}

	// Iterate through test cases
	for _, tt := range tests {
		// Run subtest for each test case
		t.Run(tt.name, func(t *testing.T) {
			// Call the function under test
			err := SendSmsAlert(tt.order)

			// Check if the error matches the expectation
			if (err != nil) != tt.wantErr {
				t.Errorf("SendSmsAlert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
