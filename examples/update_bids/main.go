package main

import (
	"context"
	"log"
	"os"

	sdk "github.com/LittleAksMax/amazon-ads-api-sdk-go"
	"github.com/LittleAksMax/amazon-ads-api-sdk-go/models"
	"github.com/joho/godotenv"
)

const profileId int64 = 867051359340884
const adGroupId string = "309325939197546"

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx := context.Background()

	refreshToken := os.Getenv("REFRESH_TOKEN")
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	redirectURI := os.Getenv("REDIRECT_URI")

	// Create config for Amazon Ads API
	authConfig := sdk.NewAmazonAuthAPIConfig(clientID, clientSecret, redirectURI)
	authClient, err := sdk.NewAmazonAuthClient(authConfig, sdk.AmazonRegions.Europe)
	if err != nil {
		log.Fatal("Error initialising Auth Client")
	}

	// Use config to create Amazon Ads API client
	client, err := sdk.NewAmazonAdsAPIClient(&sdk.Configuration{
		AuthClient: authClient,
		Region:     sdk.AmazonRegions.Europe,
	})
	if err != nil {
		log.Fatalf("Error creating Amazon Ads API client: %v", err)
	}

	client.SetRefreshToken(refreshToken)

	adGroups, err := client.AdGroupsService.GetAdGroups(profileId, &models.ListAdGroupsOptions{
		AdProductFilter: models.Filter[models.AdProduct]{
			Include: []models.AdProduct{models.AdProductSP},
		},
		AdGroupIDFilter: &models.Filter[string]{
			Include: []string{adGroupId},
		},
	}).Collect(ctx)
	if err != nil {
		log.Fatalf("Error getting ad groups: %v", err)
	}

	if len(adGroups) != 1 {
		log.Fatalf("Expected 1 ad group, got %d", len(adGroups))
	}

	adGroup := adGroups[0]
	targets, err := client.TargetsService.GetTargets(profileId, &models.ListTargetsOptions{
		AdProductFilter: models.Filter[models.AdProduct]{
			Include: []models.AdProduct{models.AdProductSP},
		},
		AdGroupIDFilter: &models.Filter[string]{
			Include: []string{adGroupId},
		},
	}).Collect(ctx)
	if err != nil {
		log.Fatalf("Error getting targets: %v", err)
	}

	_ = targets
	log.Printf("Default Bid: %f", adGroup.Bid.DefaultBid)
}
