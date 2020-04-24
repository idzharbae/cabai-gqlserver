package data

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/idzharbae/cabai-gqlserver/globalconstant"
	"github.com/idzharbae/marketplace-backend/svc/auth/authproto"
	"strconv"
	"time"
)

type User struct {
	ID            string
	Name          string
	UserName      string
	Email         string
	Phone         string
	Password      string
	Type          int32
	PhotoURL      string
	City          string
	Province      string
	AddressDetail string
	ZipCode       int32
	Description   string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Saldo         string
}

func UserFromToken(token string) (User, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(globalconstant.SECRET_KEY), nil
	})

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		user := User{
			ID:       strconv.Itoa(int(claims["id"].(float64))),
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

func UserFromProto(in *authproto.User) *User {
	return &User{
		ID:            strconv.Itoa(int(in.GetId())),
		Name:          in.GetName(),
		UserName:      in.GetUserName(),
		Email:         in.GetEmail(),
		Phone:         in.GetPhone(),
		Password:      in.GetPassword(),
		Type:          in.GetType(),
		PhotoURL:      in.GetPhotoUrl(),
		City:          in.GetCity(),
		Province:      in.GetProvince(),
		ZipCode:       in.GetZipCode(),
		AddressDetail: in.GetAddressDetail(),
		Description:   in.GetDescription(),
		CreatedAt:     time.Unix(in.GetCreatedAt(), 0),
		UpdatedAt:     time.Unix(in.GetUpdatedAt(), 0),
		Saldo:         strconv.FormatInt(in.GetSaldo(), 10),
	}
}

func (u User) GetID() int64 {
	id, _ := strconv.ParseInt(u.ID, 10, 64)
	return id
}
