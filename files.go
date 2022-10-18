package swissknife

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// IsFileExists - check file exists
func IsFileExists(filepath string) bool {
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return false
	}
	return true
}

// ReadFileToString read file to bytes
func ReadFileToBytes(filepath string) ([]byte, error) {
	if !IsFileExists(filepath) {
		return nil, fmt.Errorf("file %q not found", filepath)
	}

	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, errors.New("failed to read file: " + err.Error())
	}
	return data, nil
}

// ReadFileToString read file to string
func ReadFileToString(filepath string) (string, error) {
	data, err := ReadFileToBytes(filepath)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

// ReadFile read file to string
func ReadFile(filepath string) ([]byte, error) {
	fileBytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, errors.New("failed to read file: " + err.Error())
	}
	return fileBytes, err
}

// SaveStringToFile save arbitrary string to file
func SaveStringToFile(filepath string, content string) error {
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}

// ReadFileLines - read file to lines
func ReadFileLines(filePath string) ([]string, error) {
	lines := []string{}

	f, err := os.OpenFile(filePath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, errors.New("open file error: " + err.Error())
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}

			return nil, errors.New("read file line error: " + err.Error())
		}

		lines = append(lines, line)
	}
	return lines, nil
}
