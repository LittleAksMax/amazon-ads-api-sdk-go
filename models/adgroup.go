package models

const (
	AdGroupStateEnabled  = "ENABLED"
	AdGroupStateArchived = "ARCHIVED"
	AdGroupStatePaused   = "PAUSED"
)

type ListAdGroupsOptions struct {
	// Filtering
	AdProductFilter  Filter  `json:"adProductFilter"`
	CampaignIDFilter *Filter `json:"campaignIdFilter"`
	AdGroupIDFilter  *Filter `json:"adGroupIdFilter"`
	StateFilter      *Filter `json:"stateFilter"`
	NameFilter       *Filter `json:"nameFilter"`

	// Sort by "adGroupId", "adGroupName", "createTime", "updateTime", "state"
	SortOptions

	PaginationOptions
}

func (o *ListAdGroupsOptions) ToJSON() map[string]interface{} {
	return toJSONBodyOptions(o)
}

type AdGroup struct {
	AdGroupID           string                `json:"adGroupId"`
	CampaignID          string                `json:"campaignId"`
	Name                string                `json:"name"`
	State               string                `json:"state"`
	DefaultBid          float64               `json:"defaultBid"`
	Status              *AdGroupStatus        `json:"status"`
	CreationDateTime    string                `json:"creationDateTime"`
	LastUpdatedDateTime string                `json:"lastUpdatedDateTime"`
	Optimizations       *AdGroupOptimizations `json:"optimizations"`
	BiddingStrategy     string                `json:"biddingStrategy"`
	Tags                []string              `json:"tags"`
}

type AdGroupStatus struct {
	DeliveryStatus  string   `json:"deliveryStatus"`
	DeliveryReasons []string `json:"deliveryReasons"`
}

type AdGroupOptimizations struct {
	BidSettings *AdGroupBidSettings `json:"bidSettings"`
}

type AdGroupBidSettings struct {
	BidStrategy    string                 `json:"bidStrategy"`
	BidAdjustments *AdGroupBidAdjustments `json:"bidAdjustments"`
}

type AdGroupBidAdjustments struct {
	PlacementBidAdjustments []PlacementBidAdjustment `json:"placementBidAdjustments"`
}
