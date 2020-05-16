package requests

import upload "github.com/smithaitufe/go-graphql-upload"

type ListReview struct {
	ProductID string
	ShopID    string
	Page      int32
	Limit     int32
}
type CreateReview struct {
	UserID    string
	ProductID string
	ShopID    string
	Title     string
	Content   string
	Photo     *upload.GraphQLUpload
	Rating    float64
}
type UpdateReview struct {
	ID        string
	UserID    string
	ProductID string
	ShopID    string
	Title     string
	Content   string
	Photo     *upload.GraphQLUpload
	Rating    float64
}
type DeleteReview struct {
	ID     string
	UserID string
}
