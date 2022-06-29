package swissknife

import (
	"encoding/json"
)

// SaveStructToJSONFile save arbitrary structure to JSON file
func SaveStructToJSONFile(i interface{}, filepath string) error {
	jsonr, err := json.Marshal(i)
	if err != nil {
		return err
	}

	return SaveStringToFile(filepath, string(jsonr))
}
