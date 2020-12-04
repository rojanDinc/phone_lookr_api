package util

import (
	"io/ioutil"
	"os"
	"strings"
)

type DotEnv struct {
	envVariables map[string]string
}

func NewDotEnv() *DotEnv {
	data, err := ioutil.ReadFile(".env")
	if err != nil {
		return &DotEnv{
			envVariables: nil,
		}
	}
	dataStr := string(data)
	variables := parseEnvString(dataStr)

	envFile := &DotEnv{
		envVariables: variables,
	}

	return envFile
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

func (ef *DotEnv) GetVariable(key string) string {
	if ef.envVariables != nil {
		return ef.envVariables[key]
	}
	return os.Getenv(key)
}
