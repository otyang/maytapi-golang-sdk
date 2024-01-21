package maytapi

import (
	"github.com/otyang/maytapi-golang-sdk/client"
	"github.com/otyang/maytapi-golang-sdk/whatsapp"
)

// SDK encapsulates a client for interacting with the WhatsApp API.
type SDK struct {
	// Whatsapp holds a client instance for making API calls.
	Whatsapp *whatsapp.WhatsappSVC
}

// New creates a new SDK instance, configuring it with the provided credentials.

func New(debugMode bool, apiKey, privateKey, baseURL string) (*SDK, error) {
	client := client.New(debugMode, client.ConfigBaseURL, client.ConfigProductID, client.ConfigToken)
	return &SDK{Whatsapp: whatsapp.New(client)}, nil
}
