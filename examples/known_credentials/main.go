package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	sdk "github.com/LittleAksMax/amazon-ads-api-sdk-go"
	"github.com/joho/godotenv"
)

type RefreshToken struct {
	refreshToken string
}

func (rft *RefreshToken) Get() string {
	if rft.refreshToken != "" {
		return rft.refreshToken
	}

	client := &http.Client{}
	uri := os.Getenv("REFRESH_TOKEN_URI")
	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return ""
	}

	req.Header.Set("X-Api-Key", os.Getenv("REFRESH_TOKEN_URI_ACCESS_KEY"))
	res, err := client.Do(req)
	if err != nil {
		return ""
	}
	if res.StatusCode != http.StatusOK {
		return ""
	}
	defer res.Body.Close()

	target := struct {
		RefreshToken string `json:"refresh_token"`
	}{}

	err = json.NewDecoder(res.Body).Decode(&target)
	if err != nil {
		return ""
	}

	rft.refreshToken = target.RefreshToken
	return rft.refreshToken
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	refreshToken := RefreshToken{}

	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	redirectURI := os.Getenv("REDIRECT_URI")

	// Create config for Amazon Ads API
	authConfig := sdk.NewAmazonAuthAPIConfig(clientID, clientSecret, redirectURI)
	authClient, err := sdk.NewAmazonAuthClient(authConfig, sdk.AmazonRegions.Europe)
	if err != nil {
		log.Fatal("Error initializing Auth Client")
	}

	// Use config to create Amazon Ads API client
	client, err := sdk.NewAmazonAdsAPIClient(authClient, sdk.AmazonRegions.Europe)
	if err != nil {
		log.Fatalf("Error creating Amazon Ads API client: %v", err)
	}

	// Set refresh token (usually taken from external data source)
	log.Println(refreshToken.Get())
	client.SetRefreshToken(refreshToken.Get())

	profs, err := client.GetProfiles(context.Background(), nil)
	log.Println(len(profs))
}
