package util

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/idzharbae/cabai-gqlserver/globalconstant"
)

func UserIDFromCtx(ctx context.Context) (int64, error) {
	token, err := GetTokenFromContext(ctx)
	if err != nil {
		return 0, err
	}
	return UserIDFromToken(token)
}

func UserIDFromToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(globalconstant.SECRET_KEY), nil
	})

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		userID := int64(claims["id"].(float64))
		if err != nil {
			return 0, err
		}
		return userID, nil
	}

	return 0, err
}
