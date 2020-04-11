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
