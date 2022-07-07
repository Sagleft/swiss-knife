package swissknife

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

// GetValuePrecision - find out the precision of the numeric value
func GetValuePrecision(val float64) int {
	strs := strings.Split(strings.TrimRight(strconv.FormatFloat(val, 'f', 4, 32), "0"), ".")
	if len(strs) < 2 {
		return 0
	}
	return len(strs[1])
}

// FloatToString - convert float to string
func FloatToString(val float64) string {
	return strconv.FormatFloat(val, 'f', 8, 64)
}

// FormatDuration - format duration 52h -> "2 days 4 hours"
func FormatDuration(duration time.Duration) string {
	if duration.Seconds() < 60.0 {
		return fmt.Sprintf("%d seconds", int64(duration.Seconds()))
	}
	if duration.Minutes() < 60.0 {
		remainingSeconds := math.Mod(duration.Seconds(), 60)
		return fmt.Sprintf("%d minutes %d seconds", int64(duration.Minutes()), int64(remainingSeconds))
	}
	if duration.Hours() < 24.0 {
		remainingMinutes := math.Mod(duration.Minutes(), 60)
		remainingSeconds := math.Mod(duration.Seconds(), 60)
		return fmt.Sprintf("%d hours %d minutes %d seconds",
			int64(duration.Hours()), int64(remainingMinutes), int64(remainingSeconds))
	}
	remainingHours := math.Mod(duration.Hours(), 24)
	remainingMinutes := math.Mod(duration.Minutes(), 60)
	remainingSeconds := math.Mod(duration.Seconds(), 60)

	format := ""
	args := []interface{}{}

	daysCount := int64(duration.Hours() / 24)
	if daysCount > 0 {
		format += "%d days"
		args = append(args, daysCount)
	}
	if int64(remainingHours) > 0 {
		format += " %d hours"
		args = append(args, int64(remainingHours))
	}
	if int64(remainingMinutes) > 0 {
		format += " %d minutes"
		args = append(args, int64(remainingMinutes))
	}
	if int64(remainingSeconds) > 0 {
		format += " %d seconds"
		args = append(args, int64(remainingSeconds))
	}

	return fmt.Sprintf(format, args...)
}
