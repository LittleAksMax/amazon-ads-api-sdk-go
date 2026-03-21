package models

// BidStrategy represents the bidding strategy for campaigns
type BidStrategy string

const (
	BidStrategyAutoForSales     BidStrategy = "AUTO_FOR_SALES"
	BidStrategyManual           BidStrategy = "MANUAL"
	BidStrategyRuleBasedBidding BidStrategy = "RULE_BASED_BIDDING"
	BidStrategySalesDownOnly    BidStrategy = "SALES_DOWN_ONLY"
	BidStrategySalesUpAndDown   BidStrategy = "SALES_UP_AND_DOWN"
)

var AllBidStrategies = []BidStrategy{
	BidStrategyAutoForSales,
	BidStrategyManual,
	BidStrategyRuleBasedBidding,
	BidStrategySalesDownOnly,
	BidStrategySalesUpAndDown,
}

func (bs BidStrategy) String() string {
	return string(bs)
}

func (bs BidStrategy) IsValid() bool {
	for _, valid := range AllBidStrategies {
		if bs == valid {
			return true
		}
	}
	return false
}
