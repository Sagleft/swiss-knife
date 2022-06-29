package swissknife

// ErrorFunc - func with error result
type ErrorFunc func() error

// CheckErrors - check some errors
func CheckErrors(errChecks ...ErrorFunc) error {
	for _, errFunc := range errChecks {
		err := errFunc()
		if err != nil {
			return err
		}
	}
	return nil
}
