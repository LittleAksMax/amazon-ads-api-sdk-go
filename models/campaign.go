package models

const (
	AdProductFilterDSP = "AMAZON_DSP"
	AdProductFilterSB  = "SPONSORED_BRANDS"
	AdProductFilterSD  = "SPONSORED_DISPLAY"
	AdProductFilterSP  = "SPONSORED_PRODUCTS"
	AdProductFilterST  = "SPONSORED_TELEVISION"
)

type ListCampaignsOptions struct {
	// Filtering
	AdProductFilter   Filter  `json:"adProductFilter"`
	CampaignIDFilter  *Filter `json:"campaignIdFilter"`
	StateFilter       *Filter `json:"stateFilter"`
	NameFilter        *Filter `json:"nameFilter"`
	PortfolioIDFilter *Filter `json:"portfolioIdFilter"`

	// Sorting
	SortOptions // "campaignId", "campaignName", "createTime", "updateTime", "budget", "state"

	PaginationOptions
}

func (o *ListCampaignsOptions) ToJSON() map[string]interface{} {
	return toJSONBodyOptions(o)
}

// https://advertising.amazon.com/API/docs/en-us/amazon-ads/1-0/data-types/Campaign
type Campaign struct {
	CampaignID              string                 `json:"campaignId"`
	Name                    string                 `json:"name"`
	AdProduct               string                 `json:"adProduct"`
	State                   string                 `json:"state"`
	Status                  *CampaignStatus        `json:"status"`
	Budgets                 []Budget               `json:"budgets"`
	StartDateTime           string                 `json:"startDateTime"`
	Optimizations           *CampaignOptimizations `json:"optimizations"`
	AutoCreationSettings    *AutoCreationSettings  `json:"autoCreationSettings"`
	AutoScaleGlobalCampaign string                 `json:"autoScaleGlobalCampaign"`
	CreationDateTime        string                 `json:"creationDateTime"`
	LastUpdatedDateTime     string                 `json:"lastUpdatedDateTime"`
	MarketplaceScope        string                 `json:"marketplaceScope"`
	Marketplaces            []string               `json:"marketplaces"`
	Countries               []string               `json:"countries"`
	Tags                    []string               `json:"tags"`
}

type CampaignStatus struct {
	DeliveryStatus  string   `json:"deliveryStatus"`
	DeliveryReasons []string `json:"deliveryReasons"`
}

type CampaignOptimizations struct {
	BidSettings *BidSettings `json:"bidSettings"`
}

type BidSettings struct {
	BidStrategy    string          `json:"bidStrategy"`
	BidAdjustments *BidAdjustments `json:"bidAdjustments"`
}

type BidAdjustments struct {
	PlacementBidAdjustments []PlacementBidAdjustment `json:"placementBidAdjustments"`
}

type PlacementBidAdjustment struct {
	Placement  string `json:"placement"`
	Percentage int    `json:"percentage"`
}

type AutoCreationSettings struct {
	AutoCreateTargets bool `json:"autoCreateTargets"`
}
