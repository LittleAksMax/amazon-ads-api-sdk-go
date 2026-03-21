package models

// MarketplaceScope represents the scope of a marketplace configuration
type MarketplaceScope string

// Marketplace scope types
const (
	MarketplaceScopeGlobal            MarketplaceScope = "GLOBAL"
	MarketplaceScopeSingleMarketplace MarketplaceScope = "SINGLE_MARKETPLACE"
)

// AllMarketplaceScopes returns a slice of all valid marketplace scopes
var AllMarketplaceScopes = []MarketplaceScope{
	MarketplaceScopeGlobal,
	MarketplaceScopeSingleMarketplace,
}

// String returns the string representation of the marketplace scope
func (ms MarketplaceScope) String() string {
	return string(ms)
}

// IsValid checks if a marketplace scope is valid
func (ms MarketplaceScope) IsValid() bool {
	for _, valid := range AllMarketplaceScopes {
		if ms == valid {
			return true
		}
	}
	return false
}
