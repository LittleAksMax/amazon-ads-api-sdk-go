package models

type ListAdGroupsOptions struct {
	AdProductFilter  Filter[AdProduct] `json:"adProductFilter"`
	CampaignIDFilter *Filter[string]   `json:"campaignIdFilter"`
	AdGroupIDFilter  *Filter[string]   `json:"adGroupIdFilter"`
	StateFilter      *Filter[State]    `json:"stateFilter"`
	NameFilter       *Filter[string]   `json:"nameFilter"`

	// Sort by "adGroupId", "adGroupName", "createTime", "updateTime", "state"
	SortOptions

	PaginationOptions
}

func (o *ListAdGroupsOptions) ToJSON() map[string]interface{} {
	return toJSONBodyOptions(o)
}

type UpdateAdGroupsOptions struct {
	AdGroups []UpdateAdGroupOption `json:"adGroups"`
}

type UpdateAdGroupOption struct {
	AdGroupID string `json:"adGroupId"`

	Name                      *string                    `json:"name,omitempty"`
	State                     *string                    `json:"state,omitempty"`
	Tags                      []Tag                      `json:"tags,omitempty"`
	Bid                       *Bid                       `json:"bid,omitempty"`
	MarketplaceScope          *MarketplaceScope          `json:"marketplaceScope,omitempty"`
	Marketplaces              []Marketplace              `json:"marketplaces,omitempty"`
	MarketplaceConfigurations []AdGroupMarketplaceConfig `json:"marketplaceConfigurations,omitempty"`
}

type AdGroup struct {
	AdGroupID                    string                     `json:"adGroupId"`
	GlobalAdGroupID              string                     `json:"globalAdGroupId"`
	CampaignID                   string                     `json:"campaignId"`
	Name                         string                     `json:"name"`
	State                        State                      `json:"state"`
	AdProduct                    AdProduct                  `json:"adProduct"`
	CreationDateTime             string                     `json:"creationDateTime"`
	LastUpdatedDateTime          string                     `json:"lastUpdatedDateTime"`
	StartDateTime                string                     `json:"startDateTime"`
	EndDateTime                  string                     `json:"endDateTime"`
	Status                       *AdGroupStatus             `json:"status"`
	Bid                          *Bid                       `json:"bid"`
	Budgets                      []Budget                   `json:"budgets"`
	Optimization                 *AdGroupOptimization       `json:"optimization"`
	Pacing                       *PacingSettings            `json:"pacing"`
	MarketplaceScope             MarketplaceScope           `json:"marketplaceScope"`
	Marketplaces                 []Marketplace              `json:"marketplaces"`
	MarketplaceConfigurations    []AdGroupMarketplaceConfig `json:"marketplaceConfigurations"`
	Tags                         []Tag                      `json:"tags"`
	AdvertisedProductCategoryIds []string                   `json:"advertisedProductCategoryIds"`
	CreativeType                 CreativeType               `json:"creativeType"`
	CreativeRotationType         CreativeRotationType       `json:"creativeRotationType"`
	InventoryType                InventoryType              `json:"inventoryType"`
	PurchaseOrderNumber          string                     `json:"purchaseOrderNumber"`
	Fees                         []Fee                      `json:"fees"`
	Frequencies                  []Frequency                `json:"frequencies"`
	TargetingSettings            *TargetingSettings         `json:"targetingSettings"`
}

type AdGroupStatus struct {
	DeliveryStatus      DeliveryStatus              `json:"deliveryStatus"`
	DeliveryReasons     []DeliveryReason            `json:"deliveryReasons"`
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
	Marketplace  Marketplace `json:"marketplace"`
	DefaultBid   float64     `json:"defaultBid"`
	CurrencyCode string      `json:"currencyCode"`
}

type AdGroupOptimization struct {
	BidStrategy    BidStrategy     `json:"bidStrategy"`
	BudgetSettings *BudgetSettings `json:"budgetSettings"`
	GoalSettings   *GoalSettings   `json:"goalSettings"`
}

type BudgetSettings struct {
	BudgetAllocation   BudgetAllocation `json:"budgetAllocation"`
	DailyMinSpendValue float64          `json:"dailyMinSpendValue"`
}

type GoalSettings struct {
	KPI string `json:"kpi"`
}

type PacingSettings struct {
	DeliveryProfile DeliveryProfile `json:"deliveryProfile"`
}

type AdGroupMarketplaceConfig struct {
	AdGroupID   string                             `json:"adGroupId"`
	Marketplace Marketplace                        `json:"marketplace"`
	Overrides   *AdGroupMarketplaceConfigOverrides `json:"overrides"`
}

type AdGroupMarketplaceConfigOverrides struct {
	Name  string `json:"name"`
	State State  `json:"state"`
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

type UpdateAdGroupsResponse struct {
	Success []UpdateAdGroupSuccess `json:"success"`
	Error   []UpdateAdGroupError   `json:"error"`
}

type UpdateAdGroupSuccess struct {
	AdGroup AdGroup `json:"adGroup"`
	Index   int     `json:"index"`
}

type UpdateAdGroupError struct {
	Errors []struct {
		Code          string  `json:"code"`
		FieldLocation *string `json:"fieldLocation,omitempty"`
		Message       string  `json:"message"`
	} `json:"errors"`
	Index int `json:"index"`
}
