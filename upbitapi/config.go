package upbitapi

import (
	"context"
	"fmt"
)

type contextKey string

const ConfigKey contextKey = "credential"

type Credential struct {
	AccessKey string
	SecretKey string
}

func SetCtxCredential(ctx context.Context, cred Credential) (context.Context, error) {
	if cred.AccessKey == "" || cred.SecretKey == "" {
		return ctx, fmt.Errorf("credential is blank")
	}
	return context.WithValue(ctx, ConfigKey, cred), nil
}

func GetCtxCredential(ctx context.Context) (Credential, error) {
	if value := ctx.Value(ConfigKey); value != nil {
		if credential, ok := value.(Credential); ok {
			return credential, nil
		} else {
			return Credential{}, fmt.Errorf("credential is malformed")
		}
	}
	return Credential{}, fmt.Errorf("credential has not been set up yet")
}
