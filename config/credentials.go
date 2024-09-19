package config

import (
	"context"
	"fmt"
	"path/filepath"
)

type Credential struct {
	Tag       string `yaml:"tag"`
	AccessKey string `yaml:"accesskey"`
	SecretKey string `yaml:"secretkey"`
}

type configCredential struct {
	Credentials []Credential
}

func NewCredentials(ctx context.Context, tag string) (Credential, error) {
	if tag == "" {
		return Credential{}, fmt.Errorf("credential tag is not provided")
	}
	var cfgCred configCredential
	cfgCred = getAllConfig(defaultSharedCredentialsFilename(), cfgCred).(configCredential)
	getCredByTag := func(creds []Credential, tag string) (Credential, error) {
		for _, cred := range creds {
			if cred.Tag == tag {
				return cred, nil
			}
		}
		return Credential{}, fmt.Errorf("no credential found for tag[%s]", tag)
	}
	cred, err := getCredByTag(cfgCred.Credentials, tag)
	if err != nil {
		return getCredByTag(cfgCred.Credentials, "default")
	}
	return cred, nil
}

func defaultSharedCredentialsFilename() string {
	return filepath.Join(configPath, "credentials.yaml")
}
