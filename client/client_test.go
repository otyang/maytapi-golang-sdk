package client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeRequest_Success(t *testing.T) {
	t.Parallel()

	type APISuccessResponse struct {
		Success bool
		Status  int
		Message string
	}

	wantAPISuccessResponse := APISuccessResponse{
		Success: true,
		Status:  100,
		Message: "this is an api success message",
	}

	// Set up a mock server for testing
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(wantAPISuccessResponse)
		if err != nil {
			panic(err)
		}
	}))
	defer ts.Close()

	var (
		path                  = "/test-path"
		body                  = map[string]any{"tel": 2349011112222}
		gotAPISuccessResponse APISuccessResponse
		client                = New(true, ts.URL, "", "", "")
	)

	// 	client := New(ConfigApiKey, string(ConfigPrivateKey), ts.URL, nil)
	httpResponse, err := client.MakeRequest(context.Background(), "GET", path, body, &gotAPISuccessResponse)
	assert.NoError(t, err)
	assert.Equal(t, httpResponse.StatusCode(), http.StatusOK)
	assert.Equal(t, httpResponse.Request.Method, "GET")
	assert.Equal(t, wantAPISuccessResponse, gotAPISuccessResponse)
}

func TestMakeRequest_Error(t *testing.T) {
	t.Parallel()

	wantAPIError := APIError{
		Success: false,
		Status:  20,
		Message: "an error occured",
	}

	// Set up a mock server for testing
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode(wantAPIError)
		if err != nil {
			panic(err)
		}
	}))
	defer ts.Close()

	var (
		client  = New(true, ts.URL, "", "", "")
		path    = "/test-path"
		body    = map[string]any{"tel": 2349011112222}
		success map[string]any
	)

	httpResponse, err := client.MakeRequest(context.Background(), "post", path, body, success)
	assert.Equal(t, httpResponse.StatusCode(), http.StatusBadRequest)
	assert.Equal(t, httpResponse.Request.Method, "POST")
	assert.Error(t, err)

	// ensure the error returned was rightly parsed + correct type
	gotAPIError, ok := err.(APIError)
	assert.True(t, ok)
	assert.Equal(t, wantAPIError, gotAPIError)

	// Test INVALID-METHOD request with body
	httpResponse, err = client.MakeRequest(context.Background(), "INVALID-METHOD", path, nil, nil)
	assert.Error(t, err)
	assert.Nil(t, httpResponse)

	// Check to ensure the error not api error but normal error
	gotAPIError, ok = err.(APIError)
	assert.False(t, ok)
	assert.Equal(t, APIError{}, gotAPIError)
}

// MayTapi Live server does not really make use of http Codes
// All response codes are http.StatusOK (even when its an error)
// Al responses both success/ error uses one body structure.
// To this end its important to take note of the success field
func TestMakeRequest_Live_Server(t *testing.T) {
	t.Parallel()

	type APISuccessResponse struct {
		Success bool
		Status  int
		Message string
	}

	var (
		ctx        = context.Background()
		path       = "/s/333"
		apiSuccess APISuccessResponse
	)

	client := New(true, "https://api.maytapi.com/api", "", "", "")
	httpResponse, err := client.MakeRequest(ctx, "post", path, nil, &apiSuccess)
	assert.NoError(t, err)
	assert.NotEmpty(t, apiSuccess)
	assert.Equal(t, httpResponse.Request.Method, "POST")
}
