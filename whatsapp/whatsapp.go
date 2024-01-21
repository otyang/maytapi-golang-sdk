package whatsapp

import (
	"context"
	"fmt"
	"strings"

	"github.com/otyang/maytapi-golang-sdk/client"
)

// SendMessageParams represents the parameters required to send a WhatsApp message.
type SendMessageParams struct {
	// ToNumber is the recipient's WhatsApp phone number.
	// Must include country code without the "+"
	ToNumber string `json:"to_number"`
	// Type indicates the message type (text, forward, links etc).
	// for text use text. This is the default option
	Type string `json:"type"`
	// Message is the content of the message to be sent.
	Message string `json:"message"`
}

// SendMessageResponse represents the response received after sending a message.
type SendMessageResponse struct {
	// Success indicates whether the message was sent successfully.
	Success bool `json:"success"`
	// Data contains additional information about the sent message, if successful.
	Data struct {
		// ChatID is the unique identifier of the chat where the message was sent.
		ChatID string `json:"chatId"`
		// MsgID is the unique identifier of the sent message.
		MsgID string `json:"msgId"`
	} `json:"data"`
}

// WhatsappSVC provides methods for interacting with the WhatsApp API.
type WhatsappSVC struct {
	// client is the underlying client for making API requests.
	client *client.Client
}

func New(client *client.Client) *WhatsappSVC {
	return &WhatsappSVC{
		client: client,
	}
}

// SendMessage sends a WhatsApp message using the provided parameters.
// Reference:  https://maytapi.com/whatsapp-api-documentation
func (w *WhatsappSVC) SendMessage(ctx context.Context, p SendMessageParams) (*SendMessageResponse, error) {
	var (
		path       = fmt.Sprintf("/%s/%s/sendMessage", client.ConfigProductID, client.ConfigPhoneID)
		apiSuccess SendMessageResponse
	)

	if strings.TrimSpace(p.Type) == "" {
		p.Type = "text"
	}

	// Make the POST request to the WhatsApp API endpoint.
	if _, err := w.client.MakeRequest(ctx, "post", path, p, &apiSuccess); err != nil {
		return nil, err
	}

	// Return the successful response.
	return &apiSuccess, nil
}
