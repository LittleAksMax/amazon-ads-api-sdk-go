package models

type Filter[T any] struct {
	Include []T `json:"include,omitempty"`
}

type PaginationOptions struct {
	NextToken  string `json:"nextToken"`
	MaxResults int    `json:"maxResults"`
}

type SortOrderType string

const (
	SortOrderAsc  SortOrderType = "asc"
	SortOrderDesc SortOrderType = "desc"
)

type SortOptions struct {
	SortBy    string        `json:"sortBy"`
	SortOrder SortOrderType `json:"sortOrder"`
}

type Budget struct {
	BudgetType           BudgetType           `json:"budgetType"`
	BudgetValue          BudgetValue          `json:"budgetValue"`
	RecurrenceTimePeriod RecurrenceTimePeriod `json:"recurrenceTimePeriod"`
}

type BudgetValue struct {
	MonetaryBudgetValue MonetaryBudgetValue `json:"monetaryBudgetValue"`
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
	Marketplace     Marketplace      `json:"marketplace"`
	DeliveryStatus  DeliveryStatus   `json:"deliveryStatus"`
	DeliveryReasons []DeliveryReason `json:"deliveryReasons"`
}

type Tag struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type DateTimeFields struct {
	StartDateTime *string `json:"startDateTime"`
	EndDateTime   *string `json:"endDateTime"`
}

type MarketplaceFields struct {
	MarketplaceScope *MarketplaceScope `json:"marketplaceScope"`
	Marketplaces     []Marketplace     `json:"marketplaces"`
}
