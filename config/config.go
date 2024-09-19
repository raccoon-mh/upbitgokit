package config

import (
	"context"
	"fmt"
	"path/filepath"
)

type Env struct {
	Tag          string `yaml:"tag"`
	Method       string `yaml:"method"`
	UpbitBaseUrl string `yaml:"upbiturl"`
}

type configEnv struct {
	Environments []Env
}

func NewEnv(ctx context.Context, tag string) (Env, error) {
	if tag == "" {
		return Env{}, fmt.Errorf("env tag is not provided")
	}
	var cfgEnvs configEnv
	cfgEnvs = getAllConfig(defaultSharedConfigFilename(), cfgEnvs).(configEnv)
	getEnvByTag := func(envs []Env, tag string) (Env, error) {
		for _, cred := range envs {
			if cred.Tag == tag {
				return cred, nil
			}
		}
		return Env{}, fmt.Errorf("no env found for tag[%s]", tag)
	}
	env, err := getEnvByTag(cfgEnvs.Environments, tag)
	if err != nil {
		return getEnvByTag(cfgEnvs.Environments, "default")
	}
	return env, nil
}

func defaultSharedConfigFilename() string {
	return filepath.Join(configPath, "config.yaml")
}
