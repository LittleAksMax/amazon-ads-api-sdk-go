package models

type State string

const (
	StateEnabled  State = "ENABLED"
	StateArchived State = "ARCHIVED"
	StatePaused   State = "PAUSED"
)
