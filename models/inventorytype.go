package models

// InventoryType represents the inventory type for ad groups
type InventoryType string

const (
	InventoryTypeAAPMobileApp           InventoryType = "AAP_MOBILE_APP"
	InventoryTypeAmazonOwnedAndOperated InventoryType = "AMAZON_OWNED_AND_OPERATED"
	InventoryTypeThirdParty             InventoryType = "THIRD_PARTY"
)

var AllInventoryTypes = []InventoryType{
	InventoryTypeAAPMobileApp,
	InventoryTypeAmazonOwnedAndOperated,
	InventoryTypeThirdParty,
}

func (it InventoryType) String() string {
	return string(it)
}

func (it InventoryType) IsValid() bool {
	for _, valid := range AllInventoryTypes {
		if it == valid {
			return true
		}
	}
	return false
}
