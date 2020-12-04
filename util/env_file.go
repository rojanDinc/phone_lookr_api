package util

import (
	"io/ioutil"
	"strings"
)

type EnvFile struct {
	envVariables map[string]string
}

func NewEnvFile() (*EnvFile, error) {
	data, err := ioutil.ReadFile(".env")
	if err != nil {
		return nil, err
	}

	dataStr := string(data)
	variables := parseEnvString(dataStr)

	envFile := &EnvFile{
		envVariables: variables,
	}

	return envFile, nil
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

func (ef *EnvFile) GetVariable(key string) string {
	return ef.envVariables[key]
}
