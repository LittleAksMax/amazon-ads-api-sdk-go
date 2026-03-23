package models

type ListTargetsOptions struct {
	// Filtering
	AdProductFilter  Filter[AdProduct] `json:"adProductFilter"`
	CampaignIDFilter *Filter[string]   `json:"campaignIdFilter"`
	AdGroupIDFilter  *Filter[string]   `json:"adGroupIdFilter"`
	TargetIDFilter   *Filter[string]   `json:"targetIdFilter"`
	StateFilter      *Filter[State]    `json:"stateFilter"`

	// Sort by "targetId", "createTime", "updateTime", "state"
	SortOptions

	PaginationOptions
}

func (o *ListTargetsOptions) ToJSON() map[string]interface{} {
	return toJSONBodyOptions(o)
}

type UpdateTargetsOptions struct {
	Targets []UpdateTargetOption `json:"targets"`
}

type UpdateTargetOption struct {
	TargetID string     `json:"targetId"`
	State    *string    `json:"state,omitempty"`
	Bid      *TargetBid `json:"bid,omitempty"`
	Tags     []Tag      `json:"tags,omitempty"`
}

// https://advertising.amazon.com/API/docs/en-us/amazon-ads/1-0/data-types/Target
type Target struct {
	TargetID                  string                    `json:"targetId"`
	GlobalTargetID            string                    `json:"globalTargetId"`
	CampaignID                string                    `json:"campaignId"`
	AdGroupID                 string                    `json:"adGroupId"`
	Name                      string                    `json:"name"`
	State                     State                     `json:"state"`
	AdProduct                 AdProduct                 `json:"adProduct"`
	CreationDateTime          string                    `json:"creationDateTime"`
	LastUpdatedDateTime       string                    `json:"lastUpdatedDateTime"`
	Status                    *TargetStatus             `json:"status"`
	Bid                       *TargetBid                `json:"bid"`
	Expression                []TargetExpression        `json:"expression"`
	ResolvedExpression        []TargetExpression        `json:"resolvedExpression"`
	MarketplaceScope          MarketplaceScope          `json:"marketplaceScope"`
	Marketplaces              []Marketplace             `json:"marketplaces"`
	MarketplaceConfigurations []TargetMarketplaceConfig `json:"marketplaceConfigurations"`
	TargetDetails             TargetDetails             `json:"targetDetails"`
	Tags                      []Tag                     `json:"tags"`
}

type TargetStatus struct {
	DeliveryStatus      DeliveryStatus              `json:"deliveryStatus"`
	DeliveryReasons     []DeliveryReason            `json:"deliveryReasons"`
	MarketplaceSettings []MarketplaceDeliveryStatus `json:"marketplaceSettings"`
}

type TargetBid struct {
	Bid                 float64                       `json:"bid"`
	CurrencyCode        string                        `json:"currencyCode"`
	MarketplaceSettings []TargetBidMarketplaceSetting `json:"marketplaceSettings"`
}

type TargetBidMarketplaceSetting struct {
	Marketplace  Marketplace `json:"marketplace"`
	Bid          float64     `json:"bid"`
	CurrencyCode string      `json:"currencyCode"`
}

type TargetExpression struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type TargetMarketplaceConfig struct {
	TargetID    string                            `json:"targetId"`
	Marketplace Marketplace                       `json:"marketplace"`
	Overrides   *TargetMarketplaceConfigOverrides `json:"overrides"`
}

type TargetMarketplaceConfigOverrides struct {
	State State `json:"state"`
	Tags  []Tag `json:"tags"`
}

type UpdateTargetsResponse struct {
	Success []UpdateTargetSuccess `json:"success"`
	Error   []UpdateTargetError   `json:"error"`
}

type UpdateTargetSuccess struct {
	Target Target `json:"target"`
	Index  int    `json:"index"`
}

type UpdateTargetError struct {
	Errors []struct {
		Code          string  `json:"code"`
		FieldLocation *string `json:"fieldLocation,omitempty"`
		Message       string  `json:"message"`
	} `json:"errors"`
	Index int `json:"index"`
}
