package resolver

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/idzharbae/cabai-gqlserver/gql/auth/data"
)

type User struct {
	Data *data.User
}

func NewUser(data *data.User) *User {
	return &User{Data: data}
}

func (u *User) ID() graphql.ID {
	return graphql.ID(u.Data.ID)
}

func (u *User) Name() string {
	return u.Data.Name
}

func (u *User) UserName() string {
	return u.Data.UserName
}

func (u *User) Email() string {
	return u.Data.Email
}

func (u *User) Type() int32 {
	return u.Data.Type
}

func (u *User) Phone() string {
	return u.Data.Phone
}

func (u *User) PhotoURL() string {
	return u.Data.PhotoURL
}

func (u *User) City() string {
	return u.Data.City
}

func (u *User) Province() string {
	return u.Data.Province
}

func (u *User) AddressDetail() string {
	return u.Data.AddressDetail
}

func (u *User) ZipCode() int32 {
	return u.Data.ZipCode
}

func (u *User) Description() string {
	return u.Data.Description
}

func (u *User) CreatedAt() string {
	return u.Data.CreatedAt.Format("2006-01-02 15:04:05")
}

func (u *User) UpdatedAt() string {
	return u.Data.UpdatedAt.Format("2006-01-02 15:04:05")
}

func (u *User) Saldo() string {
	return u.Data.Saldo
}
func (u *User) TotalProducts() int32 {
	return u.Data.TotalProducts
}
func (u *User) AverageRating() float64 {
	return u.Data.AverageRating
}
