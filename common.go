package amazon_ads_api_go_sdk

// APIError represents an API error response
type APIError struct {
	Status     string
	StatusCode int
}

func (e *APIError) Error() string {
	return e.Status
}

func newAPIError(status string, statusCode int) *APIError {
	return &APIError{
		Status:     status,
		StatusCode: statusCode,
	}
}
