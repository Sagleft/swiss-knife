package swissknife

import "io/ioutil"

// ReadFileToString read file to string
func ReadFileToString(filepath string) (string, error) {
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	return string(file), err
}
