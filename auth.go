package amazon_ads_api_go_sdk

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	url2 "net/url"
	"strconv"
	"sync"
	"time"
)

type AmazonAuthAPIConfig struct {
	clientID     string
	clientSecret string
	redirectURI  string
	httpClient   *http.Client
}

type AmazonAPIAuthClient struct {
	clientID     string
	clientSecret string
	redirectURI  string

	mu        sync.RWMutex
	refreshMu sync.Mutex

	httpClient        *http.Client
	regionURL         string
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
	RefreshToken(refreshToken string) (*AmazonAPITokenResponse, error)
}

func (authClient *AmazonAPIAuthClient) RefreshToken(token string) (*AmazonAPITokenResponse, error) {
	if token == "" {
		return nil, errors.New("refresh token is empty")
	}

	queryValues := url2.Values{
		"client_id":     []string{authClient.clientID},
		"client_secret": []string{authClient.clientSecret},
		"refresh_token": []string{token},
		"grant_type":    []string{"refresh_token"},
	}

	url := url2.URL{
		Scheme:   "https",
		Host:     authClient.getRegionURL(),
		Path:     "auth/o2/token",
		RawQuery: queryValues.Encode(),
	}

	req, err := http.NewRequest(http.MethodPost, url.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")

	res, err := authClient.getHTTPClient().Do(req)
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

// generateRefreshToken exchanges an authorisation code for an access token and refresh token.
// The code is obtained from the Login with Amazon (LWA) OAuth flow redirect.
func (authClient *AmazonAPIAuthClient) generateRefreshToken(code string) (*AmazonAPITokenResponse, error) {
	queryValues := url2.Values{
		"client_id":     []string{authClient.clientID},
		"client_secret": []string{authClient.clientSecret},
		"code":          []string{code},
		"redirect_uri":  []string{authClient.redirectURI},
		"grant_type":    []string{"authorization_code"},
	}

	url := url2.URL{
		Scheme:   "https",
		Host:     authClient.getRegionURL(),
		Path:     "auth/o2/token",
		RawQuery: queryValues.Encode(),
	}

	req, err := http.NewRequest(http.MethodPost, url.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")

	res, err := authClient.getHTTPClient().Do(req)
	defer func(res *http.Response) {
		if res != nil {
			_ = res.Body.Close()
		}
	}(res)

	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		log.Println(res.StatusCode, res.Body)
		return nil, errors.New("got status code " + strconv.Itoa(res.StatusCode) + " when exchanging authorization code")
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

// ExchangeAuthorisationCode exchanges an authorisation code for tokens, stores them
// on the auth client, and returns the token response so the caller can persist the
// refresh token (e.g., in a database).
func (authClient *AmazonAPIAuthClient) ExchangeAuthorisationCode(code string) (*AmazonAPITokenResponse, error) {
	tok, err := authClient.generateRefreshToken(code)
	if err != nil {
		return nil, err
	}
	authClient.SetAccessCredentials(tok)
	return tok, nil
}

// SetRefreshToken sets the refresh token
func (authClient *AmazonAPIAuthClient) SetRefreshToken(refreshToken string) {
	authClient.mu.Lock()
	defer authClient.mu.Unlock()

	authClient.refreshTokenValue = refreshToken
}

// IsAccessTokenValid checks if the access token is still valid
func (authClient *AmazonAPIAuthClient) IsAccessTokenValid() bool {
	authClient.mu.RLock()
	defer authClient.mu.RUnlock()

	return authClient.accessToken != "" && time.Now().UTC().Before(authClient.expiresAt.UTC())
}

// SetAccessCredentials stores the token response
func (authClient *AmazonAPIAuthClient) SetAccessCredentials(tok *AmazonAPITokenResponse) {
	authClient.mu.Lock()
	defer authClient.mu.Unlock()

	authClient.accessToken = tok.AccessToken
	if tok.RefreshToken != "" {
		authClient.refreshTokenValue = tok.RefreshToken
	}
	authClient.expiresAt = time.Now().Add(time.Duration(tok.ExpiresIn) * time.Second)
}

// SetToken refreshes the access token if needed
func (authClient *AmazonAPIAuthClient) SetToken() error {
	_, err := authClient.EnsureAccessToken()
	return err
}

// EnsureAccessToken refreshes the access token if needed and returns the token used for requests.
func (authClient *AmazonAPIAuthClient) EnsureAccessToken() (string, error) {
	if token, ok := authClient.currentAccessToken(); ok {
		return token, nil
	}

	authClient.refreshMu.Lock()
	defer authClient.refreshMu.Unlock()

	if token, ok := authClient.currentAccessToken(); ok {
		return token, nil
	}

	tok, err := authClient.RefreshToken(authClient.getRefreshToken())
	if err != nil {
		return "", err
	}
	authClient.SetAccessCredentials(tok)

	return authClient.GetAccessToken(), nil
}

// GetAccessToken returns the current access token
func (authClient *AmazonAPIAuthClient) GetAccessToken() string {
	authClient.mu.RLock()
	defer authClient.mu.RUnlock()

	return authClient.accessToken
}

func (authClient *AmazonAPIAuthClient) SetRegionURL(regionURL string) {
	authClient.mu.Lock()
	defer authClient.mu.Unlock()

	authClient.regionURL = regionURL
}

// CloseIdleConnections closes any idle connections held by the auth HTTP client.
func (authClient *AmazonAPIAuthClient) CloseIdleConnections() {
	if httpClient := authClient.getHTTPClient(); httpClient != nil {
		httpClient.CloseIdleConnections()
	}
}

func (authClient *AmazonAPIAuthClient) setHTTPClient(client *http.Client) {
	authClient.mu.Lock()
	defer authClient.mu.Unlock()

	authClient.httpClient = client
}

func (authClient *AmazonAPIAuthClient) getHTTPClient() *http.Client {
	authClient.mu.RLock()
	defer authClient.mu.RUnlock()

	if authClient.httpClient == nil {
		return http.DefaultClient
	}

	return authClient.httpClient
}

func (authClient *AmazonAPIAuthClient) getRegionURL() string {
	authClient.mu.RLock()
	defer authClient.mu.RUnlock()

	return authClient.regionURL
}

func (authClient *AmazonAPIAuthClient) getRefreshToken() string {
	authClient.mu.RLock()
	defer authClient.mu.RUnlock()

	return authClient.refreshTokenValue
}

func (authClient *AmazonAPIAuthClient) currentAccessToken() (string, bool) {
	authClient.mu.RLock()
	defer authClient.mu.RUnlock()

	if authClient.accessToken == "" || !time.Now().UTC().Before(authClient.expiresAt.UTC()) {
		return "", false
	}

	return authClient.accessToken, true
}

func NewAmazonAuthClient(cfg *AmazonAuthAPIConfig, region string) (*AmazonAPIAuthClient, error) {
	regionURL, ok := amazonAuthApiRegionToURLMap[region]
	if !ok {
		return nil, errors.New("invalid region auth API: " + region)
	}

	if cfg.httpClient == nil {
		cfg.httpClient = &http.Client{}
	}

	return &AmazonAPIAuthClient{
		clientID:     cfg.clientID,
		clientSecret: cfg.clientSecret,
		redirectURI:  cfg.redirectURI,
		httpClient:   cfg.httpClient,
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
