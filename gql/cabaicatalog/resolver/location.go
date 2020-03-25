package resolver

import "github.com/idzharbae/cabai-gqlserver/gql/cabaicatalog/data"

type Location struct {
	Data *data.Location
}

func NewLocation(data *data.Location) *Location {
	return &Location{Data: data}
}

func (l *Location) Longitude() float64 {
	return l.Data.Longitude
}
func (l *Location) Latitude() float64 {
	return l.Data.Latitude
}
