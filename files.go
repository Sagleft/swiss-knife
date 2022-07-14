package swissknife

import (
	"bufio"
	"errors"
	"io"
	"io/ioutil"
	"os"
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
