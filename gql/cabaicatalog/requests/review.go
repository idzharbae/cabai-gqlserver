package requests

type ListReview struct {
	ProductID string
	ShopID    string
	Page      int32
	Limit     int32
}
