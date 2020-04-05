package data

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/idzharbae/cabai-gqlserver/gql/constant"
)

type User struct {
	ID       int64
	Name     string
	UserName string
	Email    string
	Phone    string
	Password string
	Type     int32
}

func UserFromToken(token string) (User, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(constant.SECRET_KEY), nil
	})

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		user := User{
			ID:       int64(claims["id"].(float64)),
			Name:     claims["full_name"].(string),
			UserName: claims["user_name"].(string),
			Email:    claims["email"].(string),
			Phone:    claims["phone"].(string),
			Type:     int32(claims["role"].(float64)),
		}
		if err != nil {
			return User{}, err
		}
		return user, nil
	}

	return User{}, err
}
