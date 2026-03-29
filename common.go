package amazon_ads_api_go_sdk

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

// APIError represents an API error response
type APIError struct {
	Status     string
	StatusCode int
	Body       string
}

func (e *APIError) Error() string {
	if e.Body != "" {
		return e.Status + ": " + e.Body
	}
	return e.Status
}

func newAPIError(status string, statusCode int, body string) *APIError {
	return &APIError{
		Status:     status,
		StatusCode: statusCode,
		Body:       body,
	}
}

// JSONBodyOptions is implemented by any list options struct that can be serialised to a JSON body.
type JSONBodyOptions interface {
	ToJSON() map[string]interface{}
}

// buildPaginatedPostRequest builds an HTTP POST request with a JSON body, injecting nextToken if present.
func buildPaginatedPostRequest(ctx context.Context, baseURL string, path string, profileID int64, options JSONBodyOptions, nextToken string, client *AmazonAdsAPIClient) (*http.Request, error) {
	u := url.URL{
		Scheme: "https",
		Host:   baseURL,
		Path:   path,
	}

	var bodyMap map[string]interface{}
	if options != nil {
		bodyMap = options.ToJSON()
	} else {
		bodyMap = make(map[string]interface{})
	}

	if nextToken != "" {
		bodyMap["nextToken"] = nextToken
	}

	requestBody, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.ContentLength = int64(len(requestBody))
	req.Body = io.NopCloser(bytes.NewReader(requestBody))

	headers := map[string]string{
		"Accept":                       "application/json",
		"Content-Type":                 "application/json",
		"Amazon-Advertising-API-Scope": strconv.FormatInt(profileID, 10),
	}

	client.setRequestHeaders(req, headers)

	return req, nil
}

// buildJSONRequest builds an HTTP request with the given method, a JSON-serialised body, and standard profile headers.
func buildJSONRequest(ctx context.Context, method string, baseURL string, path string, profileID int64, body interface{}, client *AmazonAdsAPIClient) (*http.Request, error) {
	u := url.URL{
		Scheme: "https",
		Host:   baseURL,
		Path:   path,
	}

	requestBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.ContentLength = int64(len(requestBody))
	req.Body = io.NopCloser(bytes.NewReader(requestBody))

	headers := map[string]string{
		"Accept":                       "application/json",
		"Content-Type":                 "application/json",
		"Amazon-Advertising-API-Scope": strconv.FormatInt(profileID, 10),
	}

	client.setRequestHeaders(req, headers)

	return req, nil
}

// doUpdateRequest performs a PUT request with JSON body, handling token refresh, error responses,
// and response body reading. Returns the raw response body bytes for the caller to unmarshal.
func doUpdateRequest(ctx context.Context, client *AmazonAdsAPIClient, path string, profileID int64, body interface{}) ([]byte, error) {
	err := client.setToken()
	if err != nil {
		return nil, err
	}

	if client.getAccessToken() == "" {
		return nil, errors.New("access token is empty after refresh")
	}

	req, err := buildJSONRequest(ctx, http.MethodPost, client.regionURL(), path, profileID, body, client)
	if err != nil {
		return nil, err
	}

	res, err := client.httpClient().Do(req)
	defer func() {
		if res != nil {
			_ = res.Body.Close()
		}
	}()
	if err != nil {
		return nil, err
	}

	// Successfully requests have status codes 2xx
	if res.StatusCode/100 != 2 {
		errBody, _ := io.ReadAll(res.Body)
		return nil, newAPIError(res.Status, res.StatusCode, string(errBody))
	}

	return io.ReadAll(res.Body)
}
