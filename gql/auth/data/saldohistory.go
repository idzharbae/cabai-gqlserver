package data

import (
	"github.com/idzharbae/marketplace-backend/svc/auth/authproto"
	"strconv"
	"time"
)

type SaldoHistory struct {
	ID           string
	UserID       string
	SourceID     string
	Description  string
	ChangeAmount string
	CreatedAt    string
	UpdatedAt    string
}

func SaldoHistoriesFromProtos(in []*authproto.SaldoHistory) []*SaldoHistory {
	res := make([]*SaldoHistory, len(in))
	for i, item := range in {
		res[i] = SaldoHistoryFromProto(item)
	}
	return res
}

func SaldoHistoryFromProto(in *authproto.SaldoHistory) *SaldoHistory {
	return &SaldoHistory{
		ID:           strconv.FormatInt(in.GetId(), 10),
		UserID:       strconv.FormatInt(in.GetUserId(), 10),
		SourceID:     strconv.FormatInt(in.GetSourceId(), 10),
		Description:  in.GetDescription(),
		ChangeAmount: strconv.FormatInt(in.GetChangeAmount(), 10),
		CreatedAt:    time.Unix(in.GetCreatedAt(), 0).Format(time.RFC3339),
		UpdatedAt:    time.Unix(in.GetUpdatedAt(), 0).Format(time.RFC3339),
	}
}

func (sh SaldoHistory) GetID() int64 {
	id, _ := strconv.ParseInt(sh.ID, 10, 64)
	return id
}

func (sh SaldoHistory) GetUserID() int64 {
	id, _ := strconv.ParseInt(sh.UserID, 10, 64)
	return id
}

func (sh SaldoHistory) GetSourceID() int64 {
	id, _ := strconv.ParseInt(sh.SourceID, 10, 64)
	return id
}
func (sh SaldoHistory) GetChangeAmount() int64 {
	id, _ := strconv.ParseInt(sh.ChangeAmount, 10, 64)
	return id
}
