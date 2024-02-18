package client

import (
	"context"
	"errors"
	"strings"

	"github.com/go-resty/resty/v2"
)

// Client represents a client for interacting with an API
type Client struct {
	restyClient *resty.Client
	debugMode   bool
	ProductID   string
	PhoneID     string
	Token       string
}

// New creates a new Client instance, configuring the base URL, headers, and API token.
// - BaseURL specifies the base endpoint for API requests.
// - Token represents the authentication token for API access.
// - ProductID identifies the specific WhatsApp product or service.
// - PhoneID identifies the device or account for WhatsApp usage.
func New(debugMode bool, baseURL, token, productID, phoneID string) *Client {
	// Create a new Resty client with the specified configurations.
	clnt := resty.New().
		SetBaseURL(baseURL).
		SetHeader("x-maytapi-key", token).
		SetHeader("Accept", "*/*").
		SetHeader("Content-Type", "application/json")

	// Return a Client instance using the configured Resty client.
	return &Client{
		restyClient: clnt,
		debugMode:   debugMode,
		ProductID:   productID,
		PhoneID:     phoneID,
		Token:       token,
	}
}

// MakeRequest makes an API request with the specified method, path, and body.
// It handles different HTTP methods (GET, POST, PUT, DELETE) and potential errors.
func (x *Client) MakeRequest(ctx context.Context, method string, path string, body any, apiSuccess any) (*resty.Response, error) {
	// Create a request object with debug mode enabled.
	rClient := x.restyClient.SetDebug(x.debugMode).R()

	// Variables to store errors and responses.
	var (
		err          error
		apiError     APIError
		httpResponse *resty.Response
	)

	// Perform the appropriate request based on the HTTP method.
	switch strings.ToUpper(method) {
	case "DELETE":
		httpResponse, err = rClient.SetContext(ctx).ForceContentType("application/json; charset=utf-8").
			SetBody(body).SetResult(apiSuccess).SetError(&apiError).Delete(path)
	case "GET":
		httpResponse, err = rClient.SetContext(ctx).SetBody(body).SetResult(apiSuccess).SetError(&apiError).Get(path)
	case "POST":
		httpResponse, err = rClient.SetContext(ctx).SetBody(body).SetResult(apiSuccess).SetError(&apiError).Post(path)
	case "PUT":
		httpResponse, err = rClient.SetContext(ctx).SetBody(body).SetResult(apiSuccess).SetError(&apiError).Put(path)
	default:
		httpResponse, err = nil, errors.New("undefined method")
	}

	return httpResponse, HandleError(err, apiError)
}
