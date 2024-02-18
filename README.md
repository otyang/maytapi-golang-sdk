# maytapi-golang-sdk


This SDK provides a convenient way to interact with the Maytapi WhatsApp API from your Go applications, enabling you to send and manage WhatsApp messages programmatically.

Note: That this library at this time only implements whats needed at this point. However it has been structured to be easily extensible for more addition.

## Installation

```bash
go get github.com/otyang/maytapi-golang-sdk
```

## Usage
 
```go
package main

import (
	"context"
	"fmt"

	"github.com/otyang/maytapi-golang-sdk"
	"github.com/otyang/maytapi-golang-sdk/whatsapp"
)

const (
	configPhoneID   = "44351"
	configBaseURL   = "https://api.maytapi.com/api"
	configToken     = "1e6bc9e1-bb76-4ce2-92ef-65f8ce206df5"
	configProductID = "e1002b74-352f-4440-8cda-f5609fb144c8"
)

func main() {
	maytapi := maytapi.New(true, configBaseURL, configToken, configProductID, configPhoneID)

	got, err := maytapi.Whatsapp.SendMessage(context.Background(), whatsapp.SendMessageParams{
		ToNumber: "+2349093****",
		Type:     "text",
		Message:  `Hello, Testing the whatsapp Mayfair API service. xo martell.`,
	})

	if err != nil {
		// handle err
	}

	fmt.Println(got.Success)     // success
	fmt.Println(got.Data.ChatID) // chat id
	fmt.Println(got.Data.MsgID)  // msg id
}
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
