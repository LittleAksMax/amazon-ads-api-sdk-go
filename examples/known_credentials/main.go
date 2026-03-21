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

	// Set refresh token (usually taken from external data source)
	log.Println(refreshToken.Get())
	client.SetRefreshToken(refreshToken.Get())

	profs, err := client.GetProfiles(ctx, nil)
	if err != nil {
		log.Fatalf("Error fetching profiles: %v", err)
	}
	log.Printf("Total profiles: %d\n", len(profs))

	// Group profiles by seller ID
	profilesBySeller := models.GroupProfilesBySellerID(profs)
	log.Printf("Total sellers: %d\n", len(profilesBySeller))

	// Iterate through each seller
	for sellerID, sellerProfiles := range profilesBySeller {
		log.Printf("\n=== Seller ID: %s (%d profile(s)) ===\n", sellerID, len(sellerProfiles))

		// Iterate through each profile for this seller
		for _, prof := range sellerProfiles {
			log.Printf("\n  Profile %d (%s - %s):\n", prof.ProfileID, prof.CountryCode, prof.AccountInfo.Name)

			// Example: Filter for only Sponsored Products campaigns that are enabled
			campaignOptions := &models.ListCampaignsOptions{
				AdProductFilter: models.Filter[models.AdProduct]{
					Include: []models.AdProduct{models.AdProductSP},
				},
				StateFilter: &models.Filter[models.State]{
					Include: []models.State{models.StateEnabled},
				},
			}

			camps, err := client.CampaignsService.GetCampaigns(ctx, prof.ProfileID, campaignOptions)
			if err != nil {
				log.Printf("Error fetching campaigns for profile %d: %v", prof.ProfileID, err)
				continue
			}
			log.Printf("-> %d campaigns\n", len(camps))

			for _, camp := range camps {
				if camp.AdProduct != models.AdProductSP || camp.State != models.StateEnabled || !camp.AutoCreationSettings.AutoCreateTargets {
					continue
				}

				adgroups, err := client.AdGroupsService.GetAdGroups(ctx, prof.ProfileID, &models.ListAdGroupsOptions{
					AdProductFilter: models.Filter[models.AdProduct]{
						Include: []models.AdProduct{models.AdProductSP},
					},
					CampaignIDFilter: &models.Filter[string]{
						Include: []string{camp.CampaignID},
					},
					StateFilter: &models.Filter[models.State]{
						Include: []models.State{models.StateEnabled},
					},
				})
				if err != nil {
					log.Printf("Error fetching AdGroups for campaign %s/profile %d: %v", camp.CampaignID, prof.ProfileID, err)
					continue
				}
				log.Printf("Campaign %s -> %d AdGroups\n", camp.Name, len(adgroups))

				for _, adgroup := range adgroups {
					ads, err := client.AdsService.GetAds(ctx, prof.ProfileID, &models.ListAdsOptions{
						AdProductFilter: models.Filter[models.AdProduct]{
							Include: []models.AdProduct{models.AdProductSP},
						},
						CampaignIDFilter: &models.Filter[string]{
							Include: []string{camp.CampaignID},
						},
						AdGroupIDFilter: &models.Filter[string]{
							Include: []string{adgroup.AdGroupID},
						},
						StateFilter: &models.Filter[models.State]{
							Include: []models.State{models.StateEnabled},
						},
					})
					if err != nil {
						log.Printf("Error fetching Ads for AdGroup %s/Campaign %s/profile %d: %v", adgroup.AdGroupID, camp.CampaignID, prof.ProfileID, err)
						continue
					}
					log.Printf("AdGroup %s -> %d Ads\n", adgroup.Name, len(ads))
				}
			}
		}
	}

}
