package amazon_ads_api_go_sdk

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	url2 "net/url"
	"strconv"
	"time"

	"github.com/LittleAksMax/amazon-ads-api-sdk-go/models"
)

type AmazonAdsAPIClient struct {
	regionURL    string
	authClient   *AmazonAPIAuthClient
	accessToken  string
	refreshToken string
	expiresAt    time.Time
}

func (aac *AmazonAdsAPIClient) SetRefreshToken(refreshToken string) {
	aac.refreshToken = refreshToken
}

func (aac *AmazonAdsAPIClient) isAccessTokenValid() bool {
	return aac.accessToken != "" && time.Now().UTC().Before(aac.expiresAt.UTC())
}

func (aac *AmazonAdsAPIClient) setAccessCredentials(tok *AmazonAPITokenResponse) {
	aac.accessToken = tok.AccessToken
	aac.refreshToken = tok.RefreshToken
	aac.expiresAt = time.Now().Add(time.Duration(tok.ExpiresIn) * time.Second)
}

func (aac *AmazonAdsAPIClient) setToken() error {
	// We already have a valid token
	if aac.isAccessTokenValid() {
		return nil
	}

	tok, err := aac.authClient.refreshToken(aac.refreshToken)
	if err != nil {
		return err
	}
	aac.setAccessCredentials(tok)

	return nil
}

func NewAmazonAdsAPIClient(authClient *AmazonAPIAuthClient, region string) (*AmazonAdsAPIClient, error) {
	regionURL, ok := amazonAdsApiRegionToURLMap[region]
	if !ok {
		return nil, errors.New("invalid region: " + region)
	}

	return &AmazonAdsAPIClient{
		authClient: authClient,
		regionURL:  regionURL,
	}, nil
}

// doRequest executes an HTTP request and returns the response body
func (aac *AmazonAdsAPIClient) doRequest(ctx context.Context, method, urlStr string, body []byte, headers map[string]string) ([]byte, error) {
	var bodyReader io.Reader
	if body != nil {
		bodyReader = bytes.NewReader(body)
	}

	req, err := http.NewRequestWithContext(ctx, method, urlStr, bodyReader)
	if err != nil {
		return nil, err
	}

	// Set common headers
	req.Header.Set("Authorization", "Bearer "+aac.accessToken)
	// Use direct assignment to preserve exact casing for Amazon-Ads-ClientId
	// We also send both header formats because Amazon Ads API is inconsistent in which one it accepts for client ID
	req.Header["Amazon-Ads-ClientId"] = []string{aac.authClient.clientID}
	req.Header["Amazon-Advertising-API-ClientId"] = []string{aac.authClient.clientID}

	// Set custom headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	defer func() {
		if res != nil {
			_ = res.Body.Close()
		}
	}()

	// Body may have important information, worth forwarding
	var resBody []byte = nil
	var resBodyReadError error = nil
	if res != nil {
		resBody, resBodyReadError = io.ReadAll(res.Body)
	}

	if err != nil {
		return resBody, err
	}

	if res.StatusCode != http.StatusOK {
		return resBody, errors.New(res.Status)
	}

	return resBody, resBodyReadError
}

func (aac *AmazonAdsAPIClient) GetProfiles(ctx context.Context, options *models.ListProfilesOptions) ([]models.Profile, error) {
	err := aac.setToken()
	if err != nil {
		return nil, err
	}

	url := url2.URL{
		Scheme: "https",
		Host:   aac.regionURL,
		Path:   "v2/profiles",
	}
	if options != nil {
		url.RawQuery = options.ToQuery().Encode()
	}

	headers := map[string]string{
		"Accept": "application/json",
	}

	bodyBytes, err := aac.doRequest(ctx, http.MethodGet, url.String(), nil, headers)
	if err != nil {
		return nil, err
	}

	var profiles []models.Profile
	err = json.Unmarshal(bodyBytes, &profiles)
	if err != nil {
		return nil, err
	}

	return profiles, nil
}

func (aac *AmazonAdsAPIClient) GetCampaigns(ctx context.Context, profileID int64, options *models.ListCampaignsOptions) ([]models.Campaign, error) {
	err := aac.setToken()
	if err != nil {
		return nil, err
	}

	url := url2.URL{
		Scheme: "https",
		Host:   aac.regionURL,
		Path:   "adsApi/v1/query/campaigns",
	}

	// Build request body using the generic toJSONBodyOptions
	var requestBody []byte
	if options != nil {
		bodyMap := options.ToJSON()
		requestBody, err = json.Marshal(bodyMap)
		if err != nil {
			return nil, err
		}
	} else {
		requestBody = nil
	}

	headers := map[string]string{
		"Accept":                       "application/json",
		"Content-Type":                 "application/json",
		"Amazon-Advertising-API-Scope": strconv.FormatInt(profileID, 10),
	}

	bodyBytes, err := aac.doRequest(ctx, http.MethodPost, url.String(), requestBody, headers)
	if err != nil {
		return nil, err
	}

	// The API returns campaigns wrapped in a "campaigns" key
	var response struct {
		Campaigns []models.Campaign `json:"campaigns"`
	}
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		return nil, err
	}

	return response.Campaigns, nil
}

func (aac *AmazonAdsAPIClient) GetAdGroups(ctx context.Context, profileID int64, options *models.ListAdGroupsOptions) ([]models.AdGroup, error) {
	err := aac.setToken()
	if err != nil {
		return nil, err
	}

	url := url2.URL{
		Scheme: "https",
		Host:   aac.regionURL,
		Path:   "adsApi/v1/query/adGroups",
	}

	// Build request body using the generic toJSONBodyOptions
	var requestBody []byte
	if options != nil {
		bodyMap := options.ToJSON()
		requestBody, err = json.Marshal(bodyMap)
		if err != nil {
			return nil, err
		}
	} else {
		requestBody = nil
	}

	headers := map[string]string{
		"Accept":                       "application/json",
		"Content-Type":                 "application/json",
		"Amazon-Advertising-API-Scope": strconv.FormatInt(profileID, 10),
	}

	bodyBytes, err := aac.doRequest(ctx, http.MethodPost, url.String(), requestBody, headers)
	if err != nil {
		return nil, err
	}

	// The API returns ad groups wrapped in an "adGroups" key
	var response struct {
		AdGroups []models.AdGroup `json:"adGroups"`
	}
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		return nil, err
	}

	return response.AdGroups, nil
}
