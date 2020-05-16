package data

import (
	"github.com/idzharbae/marketplace-backend/svc/catalog/catalogproto"
	"strconv"
	"time"
)

type Review struct {
	ID        string
	UserID    string
	ProductID string
	ShopID    string
	Title     string
	Content   string
	PhotoURL  string
	Rating    float64
	CreatedAt string
	UpdatedAt string
}

func ReviewFromProtos(protos []*catalogproto.Review) []*Review {
	prods := make([]*Review, len(protos))
	for i, proto := range protos {
		prods[i] = ReviewFromProto(proto)
	}
	return prods
}

func ReviewFromProto(in *catalogproto.Review) *Review {
	return &Review{
		ID:        strconv.FormatInt(in.GetId(), 10),
		UserID:    strconv.FormatInt(in.GetUserId(), 10),
		ProductID: strconv.FormatInt(in.GetProductId(), 10),
		ShopID:    strconv.FormatInt(in.GetShopId(), 10),
		Title:     in.GetTitle(),
		Content:   in.GetContent(),
		PhotoURL:  in.GetPhotoUrl(),
		Rating:    in.GetRating(),
		CreatedAt: time.Unix(in.GetCreatedAt(), 0).Format(time.RFC3339),
		UpdatedAt: time.Unix(in.GetUpdatedAt(), 0).Format(time.RFC3339),
	}
}
