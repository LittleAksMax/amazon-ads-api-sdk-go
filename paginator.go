package amazon_ads_api_go_sdk

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// PageFetcher builds and executes a request for a single page, given the nextToken.
// It is called by the Paginator each time Next() is invoked.
type PageFetcher func(ctx context.Context, nextToken string) (*http.Response, error)

// PageParser extracts the items and nextToken from the raw response body.
type PageParser[T any] func(body []byte) (items []T, nextToken string, err error)

// Paginator lazily iterates over paginated Amazon Ads API responses.
type Paginator[T any] struct {
	client    *AmazonAdsAPIClient
	fetch     PageFetcher
	parse     PageParser[T]
	nextToken string
	done      bool
}

// NewPaginator creates a new lazy paginator.
func NewPaginator[T any](client *AmazonAdsAPIClient, fetch PageFetcher, parse PageParser[T]) *Paginator[T] {
	return &Paginator[T]{
		client: client,
		fetch:  fetch,
		parse:  parse,
	}
}

// HasNext returns true if there may be more pages to fetch.
func (p *Paginator[T]) HasNext() bool {
	return !p.done
}

// Next fetches the next page of results.
// When there are no more pages, HasNext() will return false after this call.
func (p *Paginator[T]) Next(ctx context.Context) ([]T, error) {
	if p.done {
		return nil, errors.New("no more pages")
	}

	err := p.client.setToken()
	if err != nil {
		return nil, err
	}

	res, err := p.fetch(ctx, p.nextToken)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = res.Body.Close()
	}()

	if res.StatusCode != http.StatusOK {
		errBody, _ := io.ReadAll(res.Body)
		return nil, newAPIError(res.Status, res.StatusCode, string(errBody))
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	items, nextToken, err := p.parse(bodyBytes)
	if err != nil {
		return nil, err
	}

	if nextToken == "" {
		p.done = true
	} else {
		p.nextToken = nextToken
	}

	return items, nil
}

// Collect drains all remaining pages into a single slice.
func (p *Paginator[T]) Collect(ctx context.Context) ([]T, error) {
	var all []T
	for p.HasNext() {
		page, err := p.Next(ctx)
		if err != nil {
			return all, err
		}
		all = append(all, page...)
	}
	return all, nil
}

// newJSONParser creates a PageParser that unmarshals a JSON response where items
// live under the given key and nextToken is a top-level "nextToken" field.
func newJSONParser[T any](itemsKey string) PageParser[T] {
	return func(body []byte) ([]T, string, error) {
		var raw map[string]json.RawMessage
		if err := json.Unmarshal(body, &raw); err != nil {
			return nil, "", err
		}

		var items []T
		if data, ok := raw[itemsKey]; ok {
			if err := json.Unmarshal(data, &items); err != nil {
				return nil, "", err
			}
		}

		var nextToken string
		if data, ok := raw["nextToken"]; ok {
			// nextToken may be a JSON string or null
			_ = json.Unmarshal(data, &nextToken)
		}

		return items, nextToken, nil
	}
}
