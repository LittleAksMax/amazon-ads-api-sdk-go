package amazon_ads_api_go_sdk

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/LittleAksMax/amazon-ads-api-sdk-go/models"
)

type AdsService service

// GetAds queries ads with optional filtering and sorting.
func (s *AdsService) GetAds(ctx context.Context, profileID int64, options *models.ListAdsOptions) ([]models.Ad, error) {
	err := s.client.setToken()
	if err != nil {
		return nil, err
	}

	var body interface{}
	if options != nil {
		body = options.ToJSON()
	}

	req, err := buildJSONRequest(ctx, http.MethodPost, s.client.cfg.regionURL, "adsApi/v1/query/ads", profileID, body, s.client)
	if err != nil {
		return nil, err
	}

	res, err := s.client.cfg.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = res.Body.Close()
	}()

	if res.StatusCode != http.StatusOK {
		return nil, newAPIError(res.Status, res.StatusCode)
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response struct {
		Ads []models.Ad `json:"ads"`
	}
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		return nil, err
	}

	return response.Ads, nil
}
