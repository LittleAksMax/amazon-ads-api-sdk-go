package amazon_ads_api_go_sdk

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/LittleAksMax/amazon-ads-api-sdk-go/models"
)

type CampaignsService service

// GetCampaigns queries campaigns with optional filtering and sorting
func (s *CampaignsService) GetCampaigns(ctx context.Context, profileID int64, options *models.ListCampaignsOptions) ([]models.Campaign, error) {
	err := s.client.setToken()
	if err != nil {
		return nil, err
	}

	u := url.URL{
		Scheme: "https",
		Host:   s.client.cfg.regionURL,
		Path:   "adsApi/v1/query/campaigns",
	}

	// Build request body
	var requestBody []byte
	if options != nil {
		bodyMap := options.ToJSON()
		requestBody, err = json.Marshal(bodyMap)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u.String(), nil)
	if err != nil {
		return nil, err
	}

	if requestBody != nil {
		req.ContentLength = int64(len(requestBody))
		req.Body = io.NopCloser(bytes.NewReader(requestBody))
	}

	headers := map[string]string{
		"Accept":                       "application/json",
		"Content-Type":                 "application/json",
		"Amazon-Advertising-API-Scope": strconv.FormatInt(profileID, 10),
	}

	s.client.setRequestHeaders(req, headers)

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
		Campaigns []models.Campaign `json:"campaigns"`
	}
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		return nil, err
	}

	return response.Campaigns, nil
}
