package upbitws

import (
	"context"
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func generateJwtToken(ctx context.Context) (string, error) {
	cred, err := GetCtxCredential(ctx)
	if err != nil {
		return "", fmt.Errorf("error while GetCtxCredentialcfg : %s", err.Error())
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"access_key": cred.AccessKey,
		"nonce":      uuid.New().String(),
	})

	signedToken, err := token.SignedString([]byte(cred.SecretKey))
	if err != nil {
		return "", fmt.Errorf("error while signedString : %s", err.Error())
	}
	return signedToken, nil
}
