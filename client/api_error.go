package client

type APIError struct {
	Success bool   `json:"success"`
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (ae APIError) Error() string {
	return ae.Message
}

func (ae APIError) Empty() bool {
	return ae.Message == ""
}

// HandleError returns any non-nil http-related error (creating the request,
// getting the response, decoding) if any. If the decoded apiError is non-zero
// the apiError is returned. Otherwise, no errors occurred, returns nil.
func HandleError(httpError error, apiError APIError) error {
	if httpError != nil {
		return httpError
	}

	if !apiError.Empty() {
		return apiError
	}

	return nil
}
