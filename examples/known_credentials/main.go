package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	sdk "github.com/LittleAksMax/amazon-ads-api-sdk-go"
	"github.com/LittleAksMax/amazon-ads-api-sdk-go/models"
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
	defer func(res *http.Response) {
		if res != nil {
			err := res.Body.Close()
			if err != nil {
				log.Printf("failed to close response body: %v", err)
			}
		}
	}(res)
	if err != nil {
		return ""
	}
	if res.StatusCode != http.StatusOK {
		return ""
	}

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

	ctx := context.Background()

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
	client, err := sdk.NewAmazonAdsAPIClient(&sdk.Configuration{
		AuthClient: authClient,
		Region:     sdk.AmazonRegions.Europe,
	})
	if err != nil {
		log.Fatalf("Error creating Amazon Ads API client: %v", err)
	}

	// Set refresh token (usually taken from external data source)
	log.Println(refreshToken.Get())
	client.SetRefreshToken(refreshToken.Get())

	profs, err := client.GetProfiles(ctx, nil)
	log.Println(len(profs))

	for _, prof := range profs {
		// Example: Filter for only Sponsored Products campaigns that are enabled
		campaignOptions := &models.ListCampaignsOptions{
			AdProductFilter: models.Filter{
				Include: []string{models.AdProductFilterSP},
			},
			StateFilter: &models.Filter{
				Include: []string{models.StateEnabled},
			},
		}

		camps, err := client.CampaignsService.GetCampaigns(ctx, prof.ProfileID, campaignOptions)
		if err != nil {
			log.Printf("Error fetching campaigns for profile %d: %v", prof.ProfileID, err)
			continue
		}
		log.Printf("Profile: %d -> %d campaigns\n", prof.ProfileID, len(camps))

		for _, camp := range camps {
			adgroups, err := client.AdGroupsService.GetAdGroups(ctx, prof.ProfileID, &models.ListAdGroupsOptions{
				AdProductFilter: models.Filter{
					Include: []string{models.AdProductFilterSP},
				},
				CampaignIDFilter: &models.Filter{
					Include: []string{camp.CampaignID},
				},
				StateFilter: &models.Filter{
					Include: []string{models.StateEnabled},
				},
			})
			if err != nil {
				log.Printf("Error fetching AdGroups for campaign %s/profile %d: %v", camp.CampaignID, prof.ProfileID, err)
			}
			log.Printf("Profile %d -> Campaign %s -> %d AdGroups\n", prof.ProfileID, camp.CampaignID, len(adgroups))
			for _, adgroup := range adgroups {
				ads, err := client.AdsService.GetAds(ctx, prof.ProfileID, &models.ListAdsOptions{
					AdProductFilter: models.Filter{
						Include: []string{models.AdProductFilterSP},
					},
					CampaignIDFilter: &models.Filter{
						Include: []string{camp.CampaignID},
					},
					AdGroupIDFilter: &models.Filter{
						Include: []string{adgroup.AdGroupID},
					},
					StateFilter: &models.Filter{
						Include: []string{models.StateEnabled},
					},
				})
				if err != nil {
					log.Printf("Error fetching Ads for AdGroup %s/Campaign %s/profile %d: %v", adgroup.AdGroupID, camp.CampaignID, prof.ProfileID, err)
				}
				log.Printf("Profile %d -> Campaign %s -> AdGroup %s -> %d Ads\n", prof.ProfileID, camp.CampaignID, adgroup.AdGroupID, len(ads))
			}
		}
	}

}
