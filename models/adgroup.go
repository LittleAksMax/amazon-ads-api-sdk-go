package models

type ListAdGroupsOptions struct {
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
	AdGroupID                    string                     `json:"adGroupId"`
	GlobalAdGroupID              string                     `json:"globalAdGroupId"`
	CampaignID                   string                     `json:"campaignId"`
	Name                         string                     `json:"name"`
	State                        string                     `json:"state"`
	AdProduct                    string                     `json:"adProduct"`
	CreationDateTime             string                     `json:"creationDateTime"`
	LastUpdatedDateTime          string                     `json:"lastUpdatedDateTime"`
	StartDateTime                string                     `json:"startDateTime"`
	EndDateTime                  string                     `json:"endDateTime"`
	Status                       *AdGroupStatus             `json:"status"`
	Bid                          *Bid                       `json:"bid"`
	Budgets                      []Budget                   `json:"budgets"`
	Optimization                 *AdGroupOptimization       `json:"optimization"`
	Pacing                       *PacingSettings            `json:"pacing"`
	MarketplaceScope             string                     `json:"marketplaceScope"`
	Marketplaces                 []string                   `json:"marketplaces"`
	MarketplaceConfigurations    []AdGroupMarketplaceConfig `json:"marketplaceConfigurations"`
	Tags                         []Tag                      `json:"tags"`
	AdvertisedProductCategoryIds []string                   `json:"advertisedProductCategoryIds"`
	CreativeType                 string                     `json:"creativeType"`
	CreativeRotationType         string                     `json:"creativeRotationType"`
	InventoryType                string                     `json:"inventoryType"`
	PurchaseOrderNumber          string                     `json:"purchaseOrderNumber"`
	Fees                         []Fee                      `json:"fees"`
	Frequencies                  []Frequency                `json:"frequencies"`
	TargetingSettings            *TargetingSettings         `json:"targetingSettings"`
}

type AdGroupStatus struct {
	DeliveryStatus      string                      `json:"deliveryStatus"`
	DeliveryReasons     []string                    `json:"deliveryReasons"`
	MarketplaceSettings []MarketplaceDeliveryStatus `json:"marketplaceSettings"`
}

type Bid struct {
	BaseBid             float64                 `json:"baseBid"`
	DefaultBid          float64                 `json:"defaultBid"`
	MaxAverageBid       float64                 `json:"maxAverageBid"`
	CurrencyCode        string                  `json:"currencyCode"`
	MarketplaceSettings []BidMarketplaceSetting `json:"marketplaceSettings"`
}

type BidMarketplaceSetting struct {
	Marketplace  string  `json:"marketplace"`
	DefaultBid   float64 `json:"defaultBid"`
	CurrencyCode string  `json:"currencyCode"`
}

type AdGroupOptimization struct {
	BidStrategy    string          `json:"bidStrategy"`
	BudgetSettings *BudgetSettings `json:"budgetSettings"`
	GoalSettings   *GoalSettings   `json:"goalSettings"`
}

type BudgetSettings struct {
	BudgetAllocation   string  `json:"budgetAllocation"`
	DailyMinSpendValue float64 `json:"dailyMinSpendValue"`
}

type GoalSettings struct {
	KPI string `json:"kpi"`
}

type PacingSettings struct {
	DeliveryProfile string `json:"deliveryProfile"`
}

type AdGroupMarketplaceConfig struct {
	AdGroupID   string                             `json:"adGroupId"`
	Marketplace string                             `json:"marketplace"`
	Overrides   *AdGroupMarketplaceConfigOverrides `json:"overrides"`
}

type AdGroupMarketplaceConfigOverrides struct {
	Name  string `json:"name"`
	State string `json:"state"`
	Tags  []Tag  `json:"tags"`
}

type Fee struct {
	FeeType                string  `json:"feeType"`
	FeeValue               float64 `json:"feeValue"`
	FeeValueType           string  `json:"feeValueType"`
	CurrencyCode           string  `json:"currencyCode"`
	ThirdPartyProvider     string  `json:"thirdPartyProvider"`
	AddToBudgetSpentAmount bool    `json:"addToBudgetSpentAmount"`
}

type Frequency struct {
	TimeUnit                  string `json:"timeUnit"`
	TimeCount                 int    `json:"timeCount"`
	EventMaxCount             int    `json:"eventMaxCount"`
	FrequencyTargetingSetting string `json:"frequencyTargetingSetting"`
}

type TargetingSettings struct {
	AutomatedTargetingTactic          string             `json:"automatedTargetingTactic"`
	DefaultAudienceTargetingMatchType string             `json:"defaultAudienceTargetingMatchType"`
	UserLocationSignal                string             `json:"userLocationSignal"`
	TimeZoneType                      string             `json:"timeZoneType"`
	EnableLanguageTargeting           bool               `json:"enableLanguageTargeting"`
	SiteLanguage                      string             `json:"siteLanguage"`
	TargetedPGDealId                  string             `json:"targetedPGDealId"`
	TacticsConvertersExclusionType    string             `json:"tacticsConvertersExclusionType"`
	VideoCompletionTier               string             `json:"videoCompletionTier"`
	AmazonViewability                 *AmazonViewability `json:"amazonViewability"`
}

type AmazonViewability struct {
	ViewabilityTier                string `json:"viewabilityTier"`
	IncludeUnmeasurableImpressions bool   `json:"includeUnmeasurableImpressions"`
}
