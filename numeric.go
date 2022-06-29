package swissknife

import (
	"strconv"
	"strings"
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
