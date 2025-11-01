package amazon_ads_api_go_sdk

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	url2 "net/url"
	"strconv"
)

type AmazonAuthAPIConfig struct {
	clientID     string
	clientSecret string
	redirectURI  string
}

type AmazonAPIAuthClient struct {
	regionURL    string
	clientID     string
	clientSecret string
	redirectURI  string
}

type AmazonAPITokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
}

type Client interface {
	refreshToken(refreshToken string) (*AmazonAPITokenResponse, error)
}

func (authClient *AmazonAPIAuthClient) refreshToken(refreshToken string) (*AmazonAPITokenResponse, error) {
	queryValues := url2.Values{
		"client_id":     []string{authClient.clientID},
		"client_secret": []string{authClient.clientSecret},
		"refresh_token": []string{refreshToken},
		"grant_type":    []string{"refresh_token"},
	}

	url := url2.URL{
		Scheme:   "https",
		Host:     authClient.regionURL,
		Path:     "auth/o2/token",
		RawQuery: queryValues.Encode(),
	}

	req, err := http.NewRequest(http.MethodPost, url.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		log.Println(res.StatusCode, res.Body)
		return nil, errors.New("got status code " + strconv.Itoa(res.StatusCode) + " when refreshing access token")
	}

	defer res.Body.Close() // Don't care about unhandled error
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	tokenResponse := &AmazonAPITokenResponse{}
	err = json.Unmarshal(body, tokenResponse)
	if err != nil {
		return nil, err
	}

	return tokenResponse, nil
}

func NewAmazonAuthClient(config *AmazonAuthAPIConfig, region string) (*AmazonAPIAuthClient, error) {
	regionURL, ok := amazonAuthApiRegionToURLMap[region]
	if !ok {
		return nil, errors.New("invalid region auth API: " + region)
	}
	return &AmazonAPIAuthClient{
		clientID:     config.clientID,
		clientSecret: config.clientSecret,
		redirectURI:  config.redirectURI,
		regionURL:    regionURL,
	}, nil
}

func NewAmazonAuthAPIConfig(clientID string, clientSecret string, redirectURI string) *AmazonAuthAPIConfig {
	return &AmazonAuthAPIConfig{
		clientID:     clientID,
		clientSecret: clientSecret,
		redirectURI:  redirectURI,
	}
}
