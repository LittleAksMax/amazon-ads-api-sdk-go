package main

import (
	"context"
	"log"
	"os"
	"strings"

	sdk "github.com/LittleAksMax/amazon-ads-api-sdk-go"
	"github.com/LittleAksMax/amazon-ads-api-sdk-go/models"
	"github.com/joho/godotenv"
)

const profileID int64 = 123456789012345
const adGroupId string = "123456789012345"

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

	adGroups, err := client.AdGroupsService.GetAdGroups(profileID, &models.ListAdGroupsOptions{
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
	targets, err := client.TargetsService.GetTargets(profileID, &models.ListTargetsOptions{
		AdProductFilter: models.Filter[models.AdProduct]{
			Include: []models.AdProduct{models.AdProductSP},
		},
		AdGroupIDFilter: &models.Filter[string]{
			Include: []string{adGroupId},
		},
		StateFilter: &models.Filter[models.State]{
			Include: []models.State{models.StateEnabled},
		},
	}).Collect(ctx)
	if err != nil {
		log.Fatalf("Error getting targets: %v", err)
	}

	newBid := 0.12

	// Try to update AdGroup
	adGroup.Bid.DefaultBid = newBid
	updatedAdGroup, err := client.AdGroupsService.UpdateAdGroups(ctx, profileID, &models.UpdateAdGroupsOptions{
		AdGroups: []models.UpdateAdGroupOption{
			{
				AdGroupID: adGroupId,
				Bid:       adGroup.Bid,
			},
		},
	})
	if err != nil {
		log.Fatalf("Error updating ad group: %v", err)
	}

	if len(updatedAdGroup.Success) != 1 && updatedAdGroup.Success[0].AdGroup.Bid.DefaultBid != newBid {
		log.Print("Didn't update ad group properly")
		return
	}

	log.Print("Successfully updated ad group")

	log.Printf("Default Bid: %f\n", adGroup.Bid.DefaultBid)
	for _, target := range targets {
		var targetInfo string
		td := target.TargetDetails
		switch {
		case td.KeywordTarget != nil:
			targetInfo = "keyword=" + td.KeywordTarget.Keyword + " match=" + string(td.KeywordTarget.MatchType)
		case td.LocationTarget != nil:
			var locs []string
			for _, l := range td.LocationTarget.Locations {
				locs = append(locs, l.LocationName+" ("+string(l.LocationType)+")")
			}
			targetInfo = "locations=" + strings.Join(locs, ",")
		case td.ProductCategoryTarget != nil:
			var exprs []string
			for _, e := range td.ProductCategoryTarget.Expression {
				exprs = append(exprs, string(e.Type)+":"+e.Value)
			}
			targetInfo = "productCategory=" + strings.Join(exprs, ",")
		case td.ProductTarget != nil:
			var exprs []string
			for _, e := range td.ProductTarget.Expression {
				exprs = append(exprs, string(e.Type)+":"+e.Value)
			}
			targetInfo = "product=" + strings.Join(exprs, ",")
		case td.ThemeTarget != nil:
			targetInfo = "theme=" + td.ThemeTarget.Theme.Value + " match=" + string(td.ThemeTarget.MatchType)
		default:
			targetInfo = "unknown target type"
		}

		if target.Bid != nil {
			log.Printf("Target [%s]: bid=%f\n", targetInfo, target.Bid.Bid)
			target.Bid.Bid = newBid
		} else {
			log.Printf("Target [%s]: no bid\n", targetInfo)
		}

	}

	updateOpts := getUpdateOpts(targets)
	if len(*updateOpts) != 0 {
		updatedTargets, err := client.TargetsService.UpdateTargets(ctx, profileID, &models.UpdateTargetsOptions{
			Targets: *updateOpts,
		})
		if err != nil {
			log.Fatalf("Error updating targets: %v", err)
		}
		if len(updatedTargets.Success) != len(targets) {
			log.Print("Didn't update all targets properly")
		}

		for i, updatedTarget := range updatedTargets.Success {
			if updatedTarget.Target.Bid == nil && targets[i].Bid != nil {
				log.Print("Didn't update target properly")
			}
			if updatedTarget.Target.Bid.Bid != newBid {
				log.Print("Didn't update target properly")
			}
		}
	} else {
		log.Print("No targets to update")
	}
}

func getUpdateOpts(targets []models.Target) *[]models.UpdateTargetOption {
	updateTargets := make([]models.UpdateTargetOption, 0, len(targets))
	for i, target := range targets {
		if target.Bid == nil {
			continue
		}
		updateTargets = append(updateTargets, models.UpdateTargetOption{
			TargetID: targets[i].TargetID,
			Bid:      &models.TargetBid{Bid: targets[i].Bid.Bid},
		})
	}

	return &updateTargets
}
