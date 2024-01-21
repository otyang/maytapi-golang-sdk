package maytapi

import (
	"context"
	"testing"

	"github.com/otyang/maytapi-golang-sdk/client"
	"github.com/otyang/maytapi-golang-sdk/whatsapp"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	client.ConfigPhoneID = "41456"
	client.ConfigBaseURL = "https://api.maytapi.com/api"
	client.ConfigToken = "test-a473a550-ecbc-455f-b06e-4f8d1cb9de7a"
	client.ConfigProductID = "test-db295204-a195-4f52-b16a-6a6079c1eeab"

	maytapi, err := New(true, client.ConfigBaseURL, client.ConfigProductID, client.ConfigToken)

	assert.NoError(t, err)

	got, err := maytapi.Whatsapp.SendMessage(context.Background(), whatsapp.SendMessageParams{
		ToNumber: "+2349093****",
		Type:     "text",
		Message:  `Hello, Testing the whatsapp Mayfair API service. xo martell.`,
	})

	assert.NoError(t, err)
	assert.NotEmpty(t, got)
	assert.True(t, got.Success)
	assert.NotEmpty(t, got.Data.ChatID)
	assert.NotEmpty(t, got.Data.MsgID)
}
