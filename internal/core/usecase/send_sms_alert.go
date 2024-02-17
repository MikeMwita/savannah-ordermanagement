package usecase

import (
	"fmt"
	"github.com/MikeMwita/africastalking-go/pkg/sms"
	"github.com/MikeMwita/savannah-ordermanagement/internal/core/models"
	"github.com/MikeMwita/savannah-ordermanagement/internal/core/repository"
)

// SendSmsAlert sends an SMS alert to a customer when an order is added
func SendSmsAlert(order models.Order) error {
	// Initialize the SMS sender with the provided API key, username, and sender ID
	client := sms.SmsSender{
		ApiKey:  "3432e5e51e098ebc001db7c2544ff23504d9c2609c83ef4e23bdcea6a7cefd85",
		ApiUser: "rangechem",
		Sender:  "RANGECHEM",
	}

	// Retrieve customer details from the database using the order's customer ID
	customerRepo := repository.NewCustomerRepo(repository.DB)
	customer, err := customerRepo.GetCustomerByID(order.CustomerID)
	if err != nil {
		return err
	}

	// Construct the SMS message with personalized information
	message := fmt.Sprintf("Hello %s, your order for %s has been placed successfully. Thank you for choosing us.", customer.Name, order.Item)

	// Set the recipient's phone number
	client.Recipients = []string{customer.Phone}

	// Set the SMS message
	client.Message = message

	// Send the SMS
	response, err := client.SendSMS()
	if err != nil {
		return err
	}

	// Log the SMS response
	fmt.Printf("SMS Response: %+v\n", response)

	// Return nil error
	return nil
}
