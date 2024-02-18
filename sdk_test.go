package maytapi

import (
	"context"
	"testing"

	"github.com/otyang/maytapi-golang-sdk/whatsapp"
	"github.com/stretchr/testify/assert"
)

const (
	configPhoneID   = "41456"
	configBaseURL   = "https://api.maytapi.com/api"
	configToken     = "test-a473a550-ecbc-455f-b06e-4f8d1cb9de7a"
	configProductID = "test-db295204-a195-4f52-b16a-6a6079c1eeab"
)

func TestNew(t *testing.T) {
	maytapi := New(true, configBaseURL, configToken, configProductID, configToken)

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
