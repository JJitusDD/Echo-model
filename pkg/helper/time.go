package helpers

import (
	"fmt"
	"time"
)

// GetStartTimeDayFromDay -
func GetStartTimeDayFromDay(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, t.Location())
}

// ParseInLocalTz -.
func ParseInLocalTz(layout, value string) time.Time {
	t, err := time.ParseInLocation(layout, value, GetLocalTz())
	if err != nil {
		panic(fmt.Sprintf("time zone error: %v", err))
	}

	return t
}

// GetLocalTz -
func GetLocalTz() *time.Location {
	timezone, err := time.LoadLocation("Asia/Ho_Chi_Minh")
	if err != nil {
		panic(fmt.Sprintf("time zone error: %v", err))
	}

	return timezone
}

func GetDiffTimeHoursFromNow(pastDate time.Time) int {
	now := time.Now()
	fmt.Println(int(now.Sub(pastDate).Hours() / 24))
	return int(now.Sub(pastDate).Hours() / 24)
}
