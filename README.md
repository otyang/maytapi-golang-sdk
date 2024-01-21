 **Here's a README.md comment for the Golang code:**

# WhatsApp SDK for Go

This SDK provides a convenient way to interact with the WhatsApp API from your Go applications, enabling you to send and manage WhatsApp messages programmatically.

Note: That this library at this time only implements whats needed at this point. However it has been structured to be easily extensible for more addition.

## Installation

```bash
go get github.com/otyang/maytapi-golang-sdk
```

## Usage

1. **Import the SDK:**

```go
import (
    "github.com/otyang/maytapi-golang-sdk"
	"github.com/otyang/maytapi-golang-sdk/client"
)
```

2. **Create an SDK instance:**

```go

	client.ConfigPhoneID = "41456"
	client.ConfigBaseURL = "https://api.maytapi.com/api"
	client.ConfigToken = "a473a550-ecbc-455f-b06e-4f8d1cb9de7a"
	client.ConfigProductID = "db295204-a195-4f52-b16a-6a6079c1eeab"

    
maytapi, err := sdk.New(true, "ConfigBaseURL", "ConfigProductID", "ConfigToken") 
```

3. **Send a message:**

```go
params := sdk.SendMessageParams{
    ToNumber: "recipient_phone_number",
    Type: "text", // default is text
    Message: "Hello from the WhatsApp SDK!",
}

response, err := maytapi.Whatsapp.SendMessage(context.Background(), params)
```

## Key Features

- **Message Sending:** Send text messages to WhatsApp users.
- **Error Handling:** Manages errors gracefully, providing informative feedback.
- **Configuration:** Allows customization of API credentials and base URL.

## Additional Notes

- Replace placeholders with your actual API credentials.
- Refer to the official documentation for more details and advanced usage: https://maytapi.com/whatsapp-api-documentation

## Contributing

Contributions are welcome! Please follow the contribution guidelines.

## License

This SDK is licensed under the MIT License.
# maytapi-golang-sdk
