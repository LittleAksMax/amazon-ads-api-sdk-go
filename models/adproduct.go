package models

// AdProduct represents an Amazon Advertising product type
type AdProduct string

// Amazon Advertising product types
const (
	AdProductDSP AdProduct = "AMAZON_DSP"           // Amazon DSP
	AdProductSB  AdProduct = "SPONSORED_BRANDS"     // Sponsored Brands
	AdProductSD  AdProduct = "SPONSORED_DISPLAY"    // Sponsored Display
	AdProductSP  AdProduct = "SPONSORED_PRODUCTS"   // Sponsored Products
	AdProductST  AdProduct = "SPONSORED_TELEVISION" // Sponsored Television
)

// AllAdProducts returns a slice of all valid ad product types
var AllAdProducts = []AdProduct{
	AdProductDSP,
	AdProductSB,
	AdProductSD,
	AdProductSP,
	AdProductST,
}

// String returns the string representation of the ad product
func (ap AdProduct) String() string {
	return string(ap)
}

// IsValid checks if an ad product type is valid
func (ap AdProduct) IsValid() bool {
	for _, valid := range AllAdProducts {
		if ap == valid {
			return true
		}
	}
	return false
}
