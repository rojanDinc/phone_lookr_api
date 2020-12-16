package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type FileReader interface {
	ReadFile() ([]byte, error)
}

type fileReaderImpl struct {
	filename string
}

func NewFileReader(filename string) FileReader {
	return fileReaderImpl{
		filename: filename,
	}
}

func (f fileReaderImpl) ReadFile() ([]byte, error) {
	return ioutil.ReadFile(f.filename)
}

func LoadDotEnvFile(fileReader FileReader) error {
	data, err := fileReader.ReadFile()
	if err != nil {
		return fmt.Errorf("Could not read env file. Reason: %v", err)
	}

	text := string(data)
	envMap := parseEnvString(text)

	for key, value := range envMap {
		if err := os.Setenv(key, value); err != nil {
			return fmt.Errorf("Could not set env key: %s. Reason: %v", key, err)
		}
	}
	return nil
}

func LoadEnvValueOrDefault(key string, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func parseEnvString(env string) map[string]string {
	lines := strings.Split(env, "\n")
	envMap := make(map[string]string)
	for _, line := range lines {
		temp := strings.Split(line, "=")
		envMap[temp[0]] = temp[1]
	}
	return envMap
}
