package models

type ListAdsOptions struct {
	// Filtering
	AdProductFilter  Filter[AdProduct] `json:"adProductFilter"`
	CampaignIDFilter *Filter[string]   `json:"campaignIdFilter"`
	AdGroupIDFilter  *Filter[string]   `json:"adGroupIdFilter"`
	AdIDFilter       *Filter[string]   `json:"adIdFilter"`
	StateFilter      *Filter[State]    `json:"stateFilter"`

	// Sort by "adId", "createTime", "updateTime", "state"
	SortOptions

	PaginationOptions
}

func (o *ListAdsOptions) ToJSON() map[string]interface{} {
	return toJSONBodyOptions(o)
}

// https://advertising.amazon.com/API/docs/en-us/amazon-ads/1-0/data-types/Ad
type Ad struct {
	AdID                      string                     `json:"adId"`
	GlobalAdID                string                     `json:"globalAdId"`
	CampaignID                string                     `json:"campaignId"`
	AdGroupID                 string                     `json:"adGroupId"`
	Name                      string                     `json:"name"`
	State                     State                      `json:"state"`
	AdProduct                 AdProduct                  `json:"adProduct"`
	AdType                    AdType                     `json:"adType"`
	CreationDateTime          string                     `json:"creationDateTime"`
	LastUpdatedDateTime       string                     `json:"lastUpdatedDateTime"`
	Status                    *AdStatus                  `json:"status"`
	Creative                  *Creative                  `json:"creative"`
	ActiveCreative            *Creative                  `json:"activeCreative"`
	MarketplaceScope          MarketplaceScope           `json:"marketplaceScope"`
	Marketplaces              []Marketplace              `json:"marketplaces"`
	MarketplaceConfigurations []MarketplaceConfiguration `json:"marketplaceConfigurations"`
	Tags                      []Tag                      `json:"tags"`
}

type AdStatus struct {
	DeliveryStatus      DeliveryStatus              `json:"deliveryStatus"`
	DeliveryReasons     []DeliveryReason            `json:"deliveryReasons"`
	MarketplaceSettings []MarketplaceDeliveryStatus `json:"marketplaceSettings"`
}

type Creative struct {
	AudioCreative *AudioCreative `json:"audioCreative"`
	VideoCreative *VideoCreative `json:"videoCreative"`
	ImageCreative *ImageCreative `json:"imageCreative"`
	TextCreative  *TextCreative  `json:"textCreative"`
}

type AudioCreative struct {
	StandardAudioSettings *StandardAudioSettings `json:"standardAudioSettings"`
}

type StandardAudioSettings struct {
	Audio                  *Asset        `json:"audio"`
	Language               string        `json:"language"`
	Products               []Product     `json:"products"`
	ImpressionTrackingUrls []TrackingUrl `json:"impressionTrackingUrls"`
}

type VideoCreative struct {
	StandardVideoSettings *StandardVideoSettings `json:"standardVideoSettings"`
}

type StandardVideoSettings struct {
	Video                  *Asset        `json:"video"`
	Products               []Product     `json:"products"`
	ImpressionTrackingUrls []TrackingUrl `json:"impressionTrackingUrls"`
}

type ImageCreative struct {
	StandardImageSettings *StandardImageSettings `json:"standardImageSettings"`
}

type StandardImageSettings struct {
	Image                  *Asset        `json:"image"`
	Products               []Product     `json:"products"`
	ImpressionTrackingUrls []TrackingUrl `json:"impressionTrackingUrls"`
}

type TextCreative struct {
	SearchTextAd *SearchTextAd `json:"searchTextAd"`
}

type SearchTextAd struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Headline    string `json:"headline"`
}

type Asset struct {
	AssetID      string `json:"assetId"`
	AssetVersion string `json:"assetVersion"`
}

type TrackingUrl struct {
	URL string `json:"url"`
}

type Product struct {
	ProductID             string               `json:"productId"`
	ProductIDType         string               `json:"productIdType"`
	ResolvedProductID     string               `json:"resolvedProductId"`
	ResolvedProductIDType string               `json:"resolvedProductIdType"`
	GlobalStoreSetting    *GlobalStoreSetting  `json:"globalStoreSetting"`
	MarketplaceSettings   []MarketplaceSetting `json:"marketplaceSettings"`
}

type GlobalStoreSetting struct {
	CatalogSourceMarketplace Marketplace `json:"catalogSourceMarketplace"`
}

type MarketplaceSetting struct {
	Marketplace Marketplace `json:"marketplace"`
	AsinValue   string      `json:"asinValue"`
}

type MarketplaceConfiguration struct {
	AdID        string                      `json:"adId"`
	Marketplace Marketplace                 `json:"marketplace"`
	Overrides   *MarketplaceConfigOverrides `json:"overrides"`
}

type MarketplaceConfigOverrides struct {
	State State `json:"state"`
	Tags  []Tag `json:"tags"`
}
