package request

type CheckoutReq struct {
	CartIDs       []string
	UserID        string
	PaymentAmount string
}
type ListOrder struct {
	UserID int32
	Status string
	Page   int32
	Limit  int32
}
