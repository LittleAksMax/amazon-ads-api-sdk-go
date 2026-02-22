package amazon_ads_api_go_sdk

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	url2 "net/url"
	"strconv"
	"time"
)

type AmazonAuthAPIConfig struct {
	clientID     string
	clientSecret string
	redirectURI  string
}

type AmazonAPIAuthClient struct {
	regionURL         string
	clientID          string
	clientSecret      string
	redirectURI       string
	accessToken       string
	refreshTokenValue string
	expiresAt         time.Time
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

func (authClient *AmazonAPIAuthClient) refreshToken(token string) (*AmazonAPITokenResponse, error) {
	queryValues := url2.Values{
		"client_id":     []string{authClient.clientID},
		"client_secret": []string{authClient.clientSecret},
		"refresh_token": []string{token},
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
	defer func(res *http.Response) {
		if res != nil {
			_ = res.Body.Close()
		}
	}(res) // Don't care about unhandled error

	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		log.Println(res.StatusCode, res.Body)
		return nil, errors.New("got status code " + strconv.Itoa(res.StatusCode) + " when refreshing access token")
	}
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

// SetRefreshToken sets the refresh token
func (authClient *AmazonAPIAuthClient) SetRefreshToken(refreshToken string) {
	authClient.refreshTokenValue = refreshToken
}

// isAccessTokenValid checks if the access token is still valid
func (authClient *AmazonAPIAuthClient) isAccessTokenValid() bool {
	return authClient.accessToken != "" && time.Now().UTC().Before(authClient.expiresAt.UTC())
}

// setAccessCredentials stores the token response
func (authClient *AmazonAPIAuthClient) setAccessCredentials(tok *AmazonAPITokenResponse) {
	authClient.accessToken = tok.AccessToken
	authClient.refreshTokenValue = tok.RefreshToken
	authClient.expiresAt = time.Now().Add(time.Duration(tok.ExpiresIn) * time.Second)
}

// setToken refreshes the access token if needed
func (authClient *AmazonAPIAuthClient) setToken() error {
	// We already have a valid token
	if authClient.isAccessTokenValid() {
		return nil
	}

	tok, err := authClient.refreshToken(authClient.refreshTokenValue)
	if err != nil {
		return err
	}
	authClient.setAccessCredentials(tok)

	return nil
}

// getAccessToken returns the current access token
func (authClient *AmazonAPIAuthClient) getAccessToken() string {
	return authClient.accessToken
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
