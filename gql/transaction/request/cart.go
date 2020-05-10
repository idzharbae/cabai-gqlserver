package request

type CreateCart struct {
	ProductId  int32
	QuantityKg float64
	UserId     int64
}
type UpdateCart struct {
	CartID        int32
	NewQuantityKG float64
	UserId        int64
}
type ListCarts struct {
	UserID int64
	Page   int32
	Limit  int32
}
