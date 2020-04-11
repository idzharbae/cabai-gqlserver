package data

import (
	"github.com/idzharbae/marketplace-backend/svc/catalog/catalogproto"
	"strconv"
	"time"
)

type Shop struct {
	ID        string
	Name      string
	Address   string
	SlugName  string
	PhotoURL  string
	Location  Location
	Products  []*Product
	CreatedAt string
	UpdatedAt string
}

type Location struct {
	Latitude  float64
	Longitude float64
}

func ShopsFromProtos(protos []*catalogproto.Shop) []*Shop {
	ss := make([]*Shop, len(protos))
	for i, proto := range protos {
		ss[i] = ShopFromProto(proto)
	}

	return ss
}

func ShopFromProto(proto *catalogproto.Shop) *Shop {
	return &Shop{
		ID:       strconv.Itoa(int(proto.GetId())),
		Name:     proto.GetName(),
		SlugName: proto.GetSlug(),
		Address:  proto.GetAddress(),
		Location: Location{
			Longitude: proto.GetLocation().GetLongitude(),
			Latitude:  proto.GetLocation().GetLatitude(),
		},
		PhotoURL:  proto.GetPhotoUrl(),
		CreatedAt: time.Unix(proto.GetCreatedAt(), 0).Format(time.RFC3339),
		UpdatedAt: time.Unix(proto.GetUpdatedAt(), 0).Format(time.RFC3339),
	}
}
