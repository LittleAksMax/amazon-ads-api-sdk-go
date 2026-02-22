package models

const (
	StateEnabled  = "ENABLED"
	StateArchived = "ARCHIVED"
	StatePaused   = "PAUSED"
)

type Filter struct {
	Include []string `json:"include,omitempty"`
}

type PaginationOptions struct {
	NextToken  string `json:"nextToken"`
	MaxResults int    `json:"maxResults"`
}

type SortOptions struct {
	SortBy    string `json:"sortBy"`    // "campaignId", "campaignName", "createTime", "updateTime", "budget", "state"
	SortOrder string `json:"sortOrder"` // "asc", "desc"
}

type Budget struct {
	BudgetType           string       `json:"budgetType"`
	BudgetValue          *BudgetValue `json:"budgetValue"`
	RecurrenceTimePeriod string       `json:"recurrenceTimePeriod"`
}

type BudgetValue struct {
	MonetaryBudgetValue *MonetaryBudgetValue `json:"monetaryBudgetValue"`
}

type MonetaryBudgetValue struct {
	MonetaryBudget      *MonetaryBudget                    `json:"monetaryBudget"`
	MarketplaceSettings []MonetaryBudgetMarketplaceSetting `json:"marketplaceSettings"`
}

type MonetaryBudget struct {
	CurrencyCode string  `json:"currencyCode"`
	Value        float64 `json:"value"`
	RuleValue    float64 `json:"ruleValue"`
}

type MonetaryBudgetMarketplaceSetting struct {
	Marketplace    string          `json:"marketplace"`
	MonetaryBudget *MonetaryBudget `json:"monetaryBudget"`
}

type MarketplaceDeliveryStatus struct {
	Marketplace     string   `json:"marketplace"`
	DeliveryStatus  string   `json:"deliveryStatus"`
	DeliveryReasons []string `json:"deliveryReasons"`
}

type Tag struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
