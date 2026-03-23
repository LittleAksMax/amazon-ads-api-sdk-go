package amazon_ads_api_go_sdk

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	url2 "net/url"

	"github.com/LittleAksMax/amazon-ads-api-sdk-go/models"
)

// AmazonAdsAPIClient manages communication with the Amazon Ads API
type AmazonAdsAPIClient struct {
	cfg    *Configuration
	common service // Reuse a single struct instead of allocating one for each service on the heap

	// API Services
	CampaignsService *CampaignsService
	AdGroupsService  *AdGroupsService
	AdsService       *AdsService
	TargetsService   *TargetsService
}

type service struct {
	client *AmazonAdsAPIClient
}

type Configuration struct {
	AuthClient *AmazonAPIAuthClient
	Region     string
	HTTPClient *http.Client
	regionURL  string
}

// NewAmazonAdsAPIClient creates a new API client
func NewAmazonAdsAPIClient(cfg *Configuration) (*AmazonAdsAPIClient, error) {
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = &http.Client{}
	}

	regionURL, ok := amazonAdsApiRegionToURLMap[cfg.Region]
	if !ok {
		return nil, errors.New("invalid region: " + cfg.Region)
	}

	c := &AmazonAdsAPIClient{
		cfg: cfg,
	}
	c.cfg.regionURL = regionURL
	c.common.client = c

	// Initialise API Services - cast common service to each service type
	c.CampaignsService = (*CampaignsService)(&c.common)
	c.AdGroupsService = (*AdGroupsService)(&c.common)
	c.AdsService = (*AdsService)(&c.common)
	c.TargetsService = (*TargetsService)(&c.common)

	return c, nil
}

func (aac *AmazonAdsAPIClient) SetRefreshToken(refreshToken string) {
	aac.cfg.AuthClient.SetRefreshToken(refreshToken)
}

func (aac *AmazonAdsAPIClient) SetRegion(region string) error {
	adsURL, ok := amazonAdsApiRegionToURLMap[region]
	if !ok {
		return errors.New("invalid region: " + region)
	}
	authURL, ok := amazonAuthApiRegionToURLMap[region]
	if !ok {
		return errors.New("invalid auth region: " + region)
	}
	aac.cfg.Region = region
	aac.cfg.regionURL = adsURL
	aac.cfg.AuthClient.regionURL = authURL
	return nil
}

func (aac *AmazonAdsAPIClient) isAccessTokenValid() bool {
	return aac.cfg.AuthClient.isAccessTokenValid()
}

func (aac *AmazonAdsAPIClient) setAccessCredentials(tok *AmazonAPITokenResponse) {
	aac.cfg.AuthClient.setAccessCredentials(tok)
}

func (aac *AmazonAdsAPIClient) setToken() error {
	return aac.cfg.AuthClient.setToken()
}

func (aac *AmazonAdsAPIClient) getAccessToken() string {
	return aac.cfg.AuthClient.getAccessToken()
}

// ExchangeAuthorisationCode exchanges a Login with Amazon (LWA) authorisation code for
// an access token and refresh token. The client is immediately ready for API calls after
// this returns. The caller should persist the returned RefreshToken for future use.
func (aac *AmazonAdsAPIClient) ExchangeAuthorisationCode(code string) (*AmazonAPITokenResponse, error) {
	return aac.cfg.AuthClient.exchangeAuthorisationCode(code)
}

func (aac *AmazonAdsAPIClient) GetProfiles(ctx context.Context, options *models.ListProfilesOptions) ([]models.Profile, error) {
	err := aac.setToken()
	if err != nil {
		return nil, err
	}

	url := url2.URL{
		Scheme: "https",
		Host:   aac.cfg.regionURL,
		Path:   "v2/profiles",
	}
	if options != nil {
		url.RawQuery = options.ToQuery().Encode()
	}

	headers := map[string]string{
		"Accept": "application/json",
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	aac.setRequestHeaders(req, headers)

	res, err := aac.cfg.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = res.Body.Close()
	}()

	if res.StatusCode != http.StatusOK {
		return nil, errors.New(res.Status)
	}

	bodyBytes, err := io.ReadAll(res.Body)
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

func (aac *AmazonAdsAPIClient) setRequestHeaders(req *http.Request, headers map[string]string) {
	authHeader := "Bearer " + aac.getAccessToken()
	req.Header.Set("Authorization", authHeader)
	req.Header["Amazon-Ads-ClientId"] = []string{aac.cfg.AuthClient.clientID}
	req.Header["Amazon-Advertising-API-ClientId"] = []string{aac.cfg.AuthClient.clientID}

	for key, value := range headers {
		req.Header.Set(key, value)
	}
}
