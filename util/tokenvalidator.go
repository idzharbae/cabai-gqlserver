package util

import (
	"context"
	"errors"
	"strings"

	"github.com/idzharbae/cabai-gqlserver/globalconstant"
)

func GetTokenFromContext(ctx context.Context) (string, error) {
	header := strings.Split(ctx.Value(globalconstant.TokenKey).(string), " ")
	if len(header) != 2 {
		return "", errors.New("invalid token")
	}
	authType, token := header[0], header[1]
	if authType != globalconstant.AuthType {
		return "", errors.New("invalid authentication type")
	}
	return token, nil
}

func ValidateToken(token string) {

}
