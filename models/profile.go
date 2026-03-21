package models

import (
	url2 "net/url"
)

/*
  - https://advertising.amazon.com/API/docs/en-us/reference/2/profiles#tag/Profiles/operation/listProfiles
    {
    "profileId": 0,
    "countryCode": "BR",
    "currencyCode": "BRL",
    "dailyBudget": 0,
    "timezone": "Africa/Cairo",
    "accountInfo": {
    "marketplaceStringId": "string",
    "id": "string",
    "type": "vendor",
    "name": "string",
    "subType": "KDP_AUTHOR",
    "validPaymentMethod": true
    }
    }
*/

// Profile /
type Profile struct {
	ProfileID    int64              `json:"profileId"`
	CountryCode  string             `json:"countryCode"`
	CurrencyCode string             `json:"currencyCode"`
	DailyBudget  float64            `json:"dailyBudget"`
	TimeZone     string             `json:"timezone"`
	AccountInfo  ProfileAccountInfo `json:"accountInfo"`
}

type ProfileAccountInfo struct {
	MarketplaceStringID string `json:"marketplaceStringId"`
	ID                  string `json:"id"`
	Type                string `json:"type"`
	Name                string `json:"name"`
	SubType             string `json:"subType"`
	ValidPaymentMethod  bool   `json:"validPaymentMethod"`
}

type ListProfilesOptions struct {
	ApiProgram               string   `query:"apiProgram"`               // "billing", "campaign", "paymentMethod", "store", "report", "account", "posts" -- default is "campaign"
	AccessLevel              string   `query:"accessLevel"`              // "view", "edit" -- default is "edit"
	ProfileTypeFilter        []string `query:"profileTypeFilter"`        // choice of "seller", "vendor", "agency" -- default is all
	ValidPaymentMethodFilter string   `query:"validPaymentMethodFilter"` // "true", "false" -- default is all
}

func (lpo *ListProfilesOptions) ToQuery() url2.Values {
	return toQueryValues(lpo)
}

// GetSellerID returns the seller/account ID for this profile
func (p *Profile) GetSellerID() string {
	return p.AccountInfo.ID
}

// GroupProfilesBySellerID bins a slice of profiles by their seller ID
// Returns a map where keys are seller IDs and values are slices of profiles
func GroupProfilesBySellerID(profiles []Profile) map[string][]Profile {
	grouped := make(map[string][]Profile)
	
	for _, profile := range profiles {
		sellerID := profile.GetSellerID()
		grouped[sellerID] = append(grouped[sellerID], profile)
	}

	return grouped
}
