package amazon_ads_api_go_sdk

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/LittleAksMax/amazon-ads-api-sdk-go/models"
)

type TargetsService service

// GetTargets returns a Paginator that lazily fetches targets page by page.
func (s *TargetsService) GetTargets(profileID int64, options *models.ListTargetsOptions) *Paginator[models.Target] {
	fetch := func(ctx context.Context, nextToken string) (*http.Response, error) {
		req, err := buildPaginatedPostRequest(ctx, s.client.cfg.regionURL, "adsApi/v1/query/targets", profileID, options, nextToken, s.client)
		if err != nil {
			return nil, err
		}

		return s.client.cfg.HTTPClient.Do(req)
	}

	return NewPaginator[models.Target](s.client, fetch, newJSONParser[models.Target]("targets"))
}

// UpdateTargets updates targets with specified properties.
func (s *TargetsService) UpdateTargets(ctx context.Context, profileID int64, options *models.UpdateTargetsOptions) (*models.UpdateTargetsResponse, error) {
	bodyBytes, err := doUpdateRequest(ctx, s.client, "/adsApi/v1/update/targets", profileID, options)
	if err != nil {
		return nil, err
	}

	var response models.UpdateTargetsResponse
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
