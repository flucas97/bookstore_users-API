package dates_utils

import "time"

const (
	timeLayout = time.RFC3339
)

// GetNow returns the actual data
func GetNow() time.Time {
	return time.Now()
}

// GetNowString returns the actual date as string
func GetNowString() string {
	return GetNow().Format(timeLayout)
}
