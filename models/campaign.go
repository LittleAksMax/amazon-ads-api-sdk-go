package models

type ListCampaignsOptions struct {
	// Filtering
	AdProductFilter   Filter[AdProduct] `json:"adProductFilter"`
	CampaignIDFilter  *Filter[string]   `json:"campaignIdFilter"`
	StateFilter       *Filter[State]    `json:"stateFilter"`
	NameFilter        *Filter[string]   `json:"nameFilter"`
	PortfolioIDFilter *Filter[string]   `json:"portfolioIdFilter"`

	// Sorting
	SortOptions // "campaignId", "campaignName", "createTime", "updateTime", "budget", "state"

	PaginationOptions
}

func (o *ListCampaignsOptions) ToJSON() map[string]interface{} {
	return toJSONBodyOptions(o)
}

type Campaign struct {
	AdProduct                 AdProduct                   `json:"adProduct"`
	AutoCreationSettings      *AutoCreationSettings       `json:"autoCreationSettings"`
	AutoScaleGlobalCampaign   AutoScaleGlobalCampaign     `json:"autoScaleGlobalCampaign"`
	BrandID                   string                      `json:"brandId"`
	Budgets                   []Budget                    `json:"budgets"`
	CampaignID                string                      `json:"campaignId"`
	GlobalCampaignID          string                      `json:"globalCampaignId"`
	PortfolioID               string                      `json:"portfolioId"`
	Name                      string                      `json:"name"`
	State                     State                       `json:"state"`
	Status                    *CampaignStatus             `json:"status"`
	StartDateTime             string                      `json:"startDateTime"`
	EndDateTime               string                      `json:"endDateTime"`
	Optimizations             *CampaignOptimizations      `json:"optimizations"`
	CreationDateTime          string                      `json:"creationDateTime"`
	LastUpdatedDateTime       string                      `json:"lastUpdatedDateTime"`
	MarketplaceScope          MarketplaceScope            `json:"marketplaceScope"`
	Marketplaces              []Marketplace               `json:"marketplaces"`
	MarketplaceConfigurations []CampaignMarketplaceConfig `json:"marketplaceConfigurations"`
	Countries                 []string                    `json:"countries"`
	Tags                      []Tag                       `json:"tags"`
}

type CampaignStatus struct {
	DeliveryStatus      DeliveryStatus              `json:"deliveryStatus"`
	DeliveryReasons     []DeliveryReason            `json:"deliveryReasons"`
	MarketplaceSettings []MarketplaceDeliveryStatus `json:"marketplaceSettings"`
}

type CampaignOptimizations struct {
	BidSettings *BidSettings `json:"bidSettings"`
}

type BidSettings struct {
	BidStrategy    BidStrategy     `json:"bidStrategy"`
	BidAdjustments *BidAdjustments `json:"bidAdjustments"`
}

type BidAdjustments struct {
	PlacementBidAdjustments []PlacementBidAdjustment `json:"placementBidAdjustments"`
}

type PlacementBidAdjustment struct {
	Placement  Placement `json:"placement"`
	Percentage int       `json:"percentage"`
}

type AutoCreationSettings struct {
	AutoCreateTargets bool `json:"autoCreateTargets"`
}

type CampaignMarketplaceConfig struct {
	CampaignID  string                              `json:"campaignId"`
	Marketplace Marketplace                         `json:"marketplace"`
	Overrides   *CampaignMarketplaceConfigOverrides `json:"overrides"`
}

type CampaignMarketplaceConfigOverrides struct {
	Name  string `json:"name"`
	State State  `json:"state"`
	Tags  []Tag  `json:"tags"`
}
