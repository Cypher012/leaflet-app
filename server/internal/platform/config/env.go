package config

import (
	"fmt"
	"os"
)

func GetEnv(str string) (string, error) {
	env := os.Getenv(str)
	if env == "" {
		return "", fmt.Errorf("%s is not set in the .env file", str)
	}
	return env, nil
}

func GetEnvs(keys ...string) (map[string]string, error) {
	envs := make(map[string]string)
	for _, key := range keys {
		val, err := GetEnv(key)
		if err != nil {
			return nil, err
		}
		envs[key] = val
	}
	return envs, nil
}
