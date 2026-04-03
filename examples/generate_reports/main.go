package main

import (
	"context"
	"log"
	"os"
	"time"

	sdk "github.com/LittleAksMax/amazon-ads-api-sdk-go"
	"github.com/LittleAksMax/amazon-ads-api-sdk-go/models"
	"github.com/joho/godotenv"
)

const profileID int64 = 123456789012345

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx := context.Background()

	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	redirectURI := os.Getenv("REDIRECT_URI")
	refreshToken := os.Getenv("REFRESH_TOKEN")

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

	request := &models.RequestReportOptions{
		Name:      "simple-sp-targeting-report",
		StartDate: models.FormatDate(time.Now().Add(-14 * 24 * time.Hour)),
		EndDate:   models.FormatDate(time.Now()),
		Configuration: models.ReportConfiguration{
			AdProduct: models.AdProductSP,
			GroupBy:   []models.ReportGroupBy{models.ReportGroupByTargeting},
			Columns: []string{
				"campaignId",
				"adGroupId",
				"targeting",
				"impressions",
				"clicks",
				"cost",
			},
			ReportTypeID: models.ReportTypeSponsoredProductsCampaigns,
			TimeUnit:     models.ReportTimeUnitSummary,
			Format:       models.ReportFormatGZIPJSON,
		},
	}

	report, err := client.ReportsService.RequestReport(ctx, profileID, request)
	if err != nil {
		log.Fatalf("Error requesting report: %v", err)
	}

	log.Printf("Requested report: %s", report.ReportID())

	for {
		details, err := report.Refresh(ctx)
		if err != nil {
			log.Fatalf("Error refreshing report status: %v", err)
		}

		log.Printf("Report status: %s", details.Status)

		if details.IsTerminal() {
			break
		}

		// Wait a bit before polling again
		time.Sleep(10 * time.Second)
	}

	generatedReport, err := report.GeneratedReport(ctx)
	if err != nil {
		log.Fatalf("Error downloading generated report: %v", err)
	}

	var rows []map[string]interface{}
	if err := generatedReport.Decode(&rows); err != nil {
		log.Fatalf("Error decoding generated report: %v", err)
	}

	log.Printf("Downloaded %d rows", len(rows))
	if len(rows) > 0 {
		log.Printf("First row: %+v", rows[0])
	}
}
