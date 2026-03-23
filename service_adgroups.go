package amazon_ads_api_go_sdk

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/LittleAksMax/amazon-ads-api-sdk-go/models"
)

type AdGroupsService service

// GetAdGroups returns a Paginator that lazily fetches ad groups page by page.
func (s *AdGroupsService) GetAdGroups(profileID int64, options *models.ListAdGroupsOptions) *Paginator[models.AdGroup] {
	fetch := func(ctx context.Context, nextToken string) (*http.Response, error) {
		req, err := buildPaginatedPostRequest(ctx, s.client.cfg.regionURL, "adsApi/v1/query/adGroups", profileID, options, nextToken, s.client)
		if err != nil {
			return nil, err
		}

		return s.client.cfg.HTTPClient.Do(req)
	}

	return NewPaginator[models.AdGroup](s.client, fetch, newJSONParser[models.AdGroup]("adGroups"))
}

// UpdateAdGroups updates ad groups with specified properties
func (s *AdGroupsService) UpdateAdGroups(ctx context.Context, profileID int64, options *models.UpdateAdGroupsOptions) (*models.UpdateAdGroupsResponse, error) {
	bodyBytes, err := doUpdateRequest(ctx, s.client, "adsApi/v1/update/adGroups", profileID, options)
	if err != nil {
		return nil, err
	}

	var response models.UpdateAdGroupsResponse
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
