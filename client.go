package amazon_ads_api_go_sdk

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
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

func (aac *AmazonAdsAPIClient) GetProfiles(ctx context.Context) ([]models.Profile, error) {
	err := aac.setToken()
	if err != nil {
		return nil, err
	}

	log.Println(aac.accessToken)
	log.Println(aac.refreshToken)

	url := "https://advertising-api.amazon.com/v2/profiles"
	body := map[string]string{
		"Authorization":                   "bearer " + aac.accessToken,
		"Amazon-Advertising-API-ClientId": aac.authClient.clientID,
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, errors.New(res.Status)
	}
	defer res.Body.Close()
	bodyBytes, err = io.ReadAll(res.Body)

	bodyString := string(bodyBytes)
	log.Println(bodyString)

	return nil, errors.New("not implemented")
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
