package config

import (
	"context"
	"fmt"
)

type contextKey string

const ConfigKey contextKey = "config"

type Config struct {
	Env        Env
	Credential Credential
}

func LoadConfig(ctx context.Context, tag string) (context.Context, error) {
	cred, err := NewCredentials(ctx, tag)
	if err != nil {
		return ctx, err
	}
	env, err := NewEnv(ctx, tag)
	if err != nil {
		return ctx, err
	}
	return context.WithValue(ctx, ConfigKey, Config{
		Credential: cred,
		Env:        env,
	}), nil
}

func GetCtxAllConfig(ctx context.Context) (Config, error) {
	if value := ctx.Value(ConfigKey); value != nil {
		if config, ok := value.(Config); ok {
			return config, nil
		} else {
			return Config{}, fmt.Errorf("config setting is not correct")
		}
	}
	return Config{}, fmt.Errorf("config has not been set up yet")
}

func GetCtxEnvConfig(ctx context.Context) (Env, error) {
	if value := ctx.Value(ConfigKey); value != nil {
		if config, ok := value.(Config); ok {
			return config.Env, nil
		} else {
			return Env{}, fmt.Errorf("config Env setting is not correct")
		}
	}
	return Env{}, fmt.Errorf("config Env has not been set up yet")
}

func GetCtxCredentialConfig(ctx context.Context) (Credential, error) {
	if value := ctx.Value(ConfigKey); value != nil {
		if config, ok := value.(Config); ok {
			return config.Credential, nil
		} else {
			return Credential{}, fmt.Errorf("config Credential setting is not correct")
		}
	}
	return Credential{}, fmt.Errorf("config Credential has not been set up yet")
}
