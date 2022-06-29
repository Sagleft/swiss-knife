package swissknife

// Ternary operator. conditional operator
// usage example: var res = ternary(val > 0, "positive", "negative")
func Ternary(statement bool, a, b interface{}) interface{} {
	if statement {
		return a
	}
	return b
}

// RunInBackground - blocking method with no exit
func RunInBackground() {
	forever := make(chan bool)
	// background work
	<-forever
}
