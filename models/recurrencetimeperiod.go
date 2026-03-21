package models

// RecurrenceTimePeriod represents the recurrence period for budgets
type RecurrenceTimePeriod string

const (
	RecurrenceTimePeriodDaily RecurrenceTimePeriod = "DAILY"
)

var AllRecurrenceTimePeriods = []RecurrenceTimePeriod{
	RecurrenceTimePeriodDaily,
}

func (rtp RecurrenceTimePeriod) String() string {
	return string(rtp)
}

func (rtp RecurrenceTimePeriod) IsValid() bool {
	for _, valid := range AllRecurrenceTimePeriods {
		if rtp == valid {
			return true
		}
	}
	return false
}
