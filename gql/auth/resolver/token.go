package resolver

import "github.com/idzharbae/cabai-gqlserver/gql/auth/data"

type Token struct {
	Data *data.Token
}

func NewToken(data *data.Token) *Token {
	return &Token{Data: data}
}

func (t *Token) Token() string {
	return t.Data.Token
}
