package models

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
