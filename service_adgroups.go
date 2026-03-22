package amazon_ads_api_go_sdk

import (
	"context"
	"encoding/json"
	"io"
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
func (s *AdGroupsService) UpdateAdGroups(ctx context.Context, profileID int64, options *models.UpdateAdGroupsOptions) ([]models.AdGroup, error) {
	err := s.client.setToken()
	if err != nil {
		return nil, err
	}

	req, err := buildJSONRequest(ctx, http.MethodPut, s.client.cfg.regionURL, "adsApi/v1/adGroups", profileID, options, s.client)
	if err != nil {
		return nil, err
	}

	res, err := s.client.cfg.HTTPClient.Do(req)
	defer func() {
		if res != nil {
			_ = res.Body.Close()
		}
	}()
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, newAPIError(res.Status, res.StatusCode)
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response struct {
		AdGroups []models.AdGroup `json:"adGroups"`
	}
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		return nil, err
	}

	return response.AdGroups, nil
}
