package amazon_ads_api_go_sdk

import (
	"context"
	"net/http"

	"github.com/LittleAksMax/amazon-ads-api-sdk-go/models"
)

type CampaignsService service

// GetCampaigns returns a Paginator that lazily fetches campaigns page by page.
func (s *CampaignsService) GetCampaigns(profileID int64, options *models.ListCampaignsOptions) *Paginator[models.Campaign] {
	fetch := func(ctx context.Context, nextToken string) (*http.Response, error) {
		req, err := buildPaginatedPostRequest(ctx, s.client.cfg.regionURL, "adsApi/v1/query/campaigns", profileID, options, nextToken, s.client)
		if err != nil {
			return nil, err
		}

		return s.client.cfg.HTTPClient.Do(req)
	}

	return NewPaginator[models.Campaign](s.client, fetch, newJSONParser[models.Campaign]("campaigns"))
}
