package maytapi

import (
	"context"
	"testing"

	"github.com/otyang/maytapi-golang-sdk/whatsapp"
	"github.com/stretchr/testify/assert"
)

const (
	configPhoneID   = "44351"
	configBaseURL   = "https://api.maytapi.com/api"
	configToken     = "1e6bc9e1-bb76-4ce2-92ef-65f8ce206df5"
	configProductID = "e1002b74-352f-4440-8cda-f5609fb144c8"
)

func TestNew(t *testing.T) {
	maytapi := New(true, configBaseURL, configToken, configProductID, configPhoneID)

	got, err := maytapi.Whatsapp.SendMessage(context.Background(), whatsapp.SendMessageParams{
		ToNumber: "+2349123456789",
		Type:     "text",
		Message:  `Hello, Testing the whatsapp Mayfair API service. xo martell.`,
	})

	assert.NoError(t, err)
	assert.NotEmpty(t, got)
	assert.True(t, got.Success)
	assert.NotEmpty(t, got.Data.ChatID)
	assert.NotEmpty(t, got.Data.MsgID)
}
