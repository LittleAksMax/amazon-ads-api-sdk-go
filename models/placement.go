package models

// Placement represents ad placement types
type Placement string

const (
	PlacementTopOfSearch  Placement = "TOP_OF_SEARCH"
	PlacementRestOfSearch Placement = "REST_OF_SEARCH"
	PlacementProductPage  Placement = "PRODUCT_PAGE"
)

var AllPlacements = []Placement{
	PlacementTopOfSearch,
	PlacementRestOfSearch,
	PlacementProductPage,
}

func (p Placement) String() string {
	return string(p)
}

func (p Placement) IsValid() bool {
	for _, valid := range AllPlacements {
		if p == valid {
			return true
		}
	}
	return false
}
