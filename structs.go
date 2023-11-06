package swissknife

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

// SaveStructToJSONFile save structure to JSON file
func SaveStructToJSONFile(i interface{}, filepath string) error {
	jsonr, err := json.Marshal(i)
	if err != nil {
		return err
	}

	return SaveStringToFile(filepath, string(jsonr))
}

// SaveStructToJSONFileIndent save structure to JSON file
func SaveStructToJSONFileIndent(i interface{}, filepath string) error {
	jsonr, err := json.MarshalIndent(i, "", "	")
	if err != nil {
		return err
	}

	return SaveStringToFile(filepath, string(jsonr))
}

// SaveStructToYamlFile save structure to YAML file (in plain text)
func SaveStructToYamlFile(i interface{}, filepath string) error {
	dataEncoded, err := yaml.Marshal(i)
	if err != nil {
		return err
	}

	return SaveStringToFile(filepath, string(dataEncoded))
}
