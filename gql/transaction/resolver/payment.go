package resolver

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/idzharbae/cabai-gqlserver/gql/transaction/data"
	"github.com/idzharbae/cabai-gqlserver/gql/transaction/enum"
	"github.com/idzharbae/cabai-gqlserver/util"

	"strconv"
)

type Payment struct {
	Data *data.Payment
}

func NewPayment(payment *data.Payment) *Payment {
	return &Payment{Data: payment}
}

func (p *Payment) ID() graphql.ID {
	return graphql.ID(p.Data.ID)
}
func (p *Payment) Amount() string {
	return strconv.FormatInt(p.Data.Amount, 10)
}
func (p *Payment) Method() string {
	if p.Data.PaymentMethod == enum.PaymentMethodSaldoCode {
		return enum.PaymentMethodSaldoString
	}
	return ""
}
func (p *Payment) Status() string {
	return util.PaymentCodeToString(int(p.Data.PaymentStatus))
}
func (p *Payment) CreatedAt() string {
	return p.Data.CreatedAt
}
func (p *Payment) UpdatedAt() string {
	return p.Data.UpdatedAt
}
