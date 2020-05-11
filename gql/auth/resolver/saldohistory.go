package resolver

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/idzharbae/cabai-gqlserver/gql/auth/data"
)

type SaldoHistory struct {
	Data *data.SaldoHistory
}

func NewSaldoHistories(ds []*data.SaldoHistory) []*SaldoHistory {
	prods := make([]*SaldoHistory, 0, len(ds))
	for _, d := range ds {
		if d == nil {
			continue
		}
		prods = append(prods, NewSaldoHistory(d))
	}
	return prods
}

func NewSaldoHistory(data *data.SaldoHistory) *SaldoHistory {
	return &SaldoHistory{Data: data}
}

func (sh SaldoHistory) ID() graphql.ID {
	return graphql.ID(sh.Data.ID)
}
func (sh SaldoHistory) UserID() graphql.ID {
	return graphql.ID(sh.Data.UserID)
}
func (sh SaldoHistory) SourceID() graphql.ID {
	return graphql.ID(sh.Data.SourceID)
}
func (sh SaldoHistory) ChangeAmount() string {
	return sh.Data.ChangeAmount
}
func (sh SaldoHistory) Description() string {
	return sh.Data.Description
}
func (sh SaldoHistory) CreatedAt() string {
	return sh.Data.CreatedAt
}
func (sh SaldoHistory) UpdatedAt() string {
	return sh.Data.UpdatedAt
}
