package swissknife

import (
	"errors"
	"io/ioutil"
)

// ReadFileToString read file to string
func ReadFileToString(filepath string) (string, error) {
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		return "", errors.New("failed to read file: " + err.Error())
	}
	return string(file), err
}

// ReadFile read file to string
func ReadFile(filepath string) ([]byte, error) {
	fileBytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, errors.New("failed to read file: " + err.Error())
	}
	return fileBytes, err
}
