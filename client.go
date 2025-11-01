package amazon_ads_api_go_sdk

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	url2 "net/url"
	"time"

	"github.com/LittleAksMax/amazon-ads-api-go-sdk/models"
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

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+aac.accessToken)
	req.Header.Set("Amazon-Advertising-API-ClientId", aac.authClient.clientID)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, errors.New(res.Status)
	}
	defer res.Body.Close()
	bodyBytes, err := io.ReadAll(res.Body)

	var profiles []models.Profile
	err = json.Unmarshal(bodyBytes, &profiles)
	if err != nil {
		return nil, err
	}

	return profiles, nil
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
