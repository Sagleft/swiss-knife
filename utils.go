package swissknife

// Ternary operator. conditional operator
// usage example: var res = ternary(val > 0, "positive", "negative")
func Ternary(statement bool, a, b interface{}) interface{} {
	if statement {
		return a
	}
	return b
}
