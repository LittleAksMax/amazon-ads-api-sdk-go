package models

// Marketplace represents an Amazon marketplace country code
type Marketplace string

// Amazon marketplace country codes
const (
	MarketplaceAE Marketplace = "AE" // United Arab Emirates
	MarketplaceAU Marketplace = "AU" // Australia
	MarketplaceBE Marketplace = "BE" // Belgium
	MarketplaceBR Marketplace = "BR" // Brazil
	MarketplaceCA Marketplace = "CA" // Canada
	MarketplaceDE Marketplace = "DE" // Germany
	MarketplaceEG Marketplace = "EG" // Egypt
	MarketplaceES Marketplace = "ES" // Spain
	MarketplaceFR Marketplace = "FR" // France
	MarketplaceGB Marketplace = "GB" // United Kingdom
	MarketplaceIE Marketplace = "IE" // Ireland
	MarketplaceIN Marketplace = "IN" // India
	MarketplaceIT Marketplace = "IT" // Italy
	MarketplaceJP Marketplace = "JP" // Japan
	MarketplaceMX Marketplace = "MX" // Mexico
	MarketplaceNL Marketplace = "NL" // Netherlands
	MarketplacePL Marketplace = "PL" // Poland
	MarketplaceSA Marketplace = "SA" // Saudi Arabia
	MarketplaceSE Marketplace = "SE" // Sweden
	MarketplaceSG Marketplace = "SG" // Singapore
	MarketplaceTR Marketplace = "TR" // Turkey
	MarketplaceUS Marketplace = "US" // United States
	MarketplaceZA Marketplace = "ZA" // South Africa
)

// AllMarketplaces returns a slice of all valid marketplace codes
var AllMarketplaces = []Marketplace{
	MarketplaceAE, MarketplaceAU, MarketplaceBE, MarketplaceBR,
	MarketplaceCA, MarketplaceDE, MarketplaceEG, MarketplaceES,
	MarketplaceFR, MarketplaceGB, MarketplaceIE, MarketplaceIN,
	MarketplaceIT, MarketplaceJP, MarketplaceMX, MarketplaceNL,
	MarketplacePL, MarketplaceSA, MarketplaceSE, MarketplaceSG,
	MarketplaceTR, MarketplaceUS, MarketplaceZA,
}

// String returns the string representation of the marketplace code
func (m Marketplace) String() string {
	return string(m)
}

// IsValid checks if a marketplace code is valid
func (m Marketplace) IsValid() bool {
	for _, valid := range AllMarketplaces {
		if m == valid {
			return true
		}
	}
	return false
}
