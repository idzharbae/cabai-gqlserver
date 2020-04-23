package request

type CheckoutReq struct {
	CartIDs       []string
	UserID        string
	PaymentAmount string
}
