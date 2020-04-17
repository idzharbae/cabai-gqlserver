package requests

// Get Token
type Login struct {
	UserNameOrEmail string
	Password        string
}

type Register struct {
	UserName      string
	Email         string
	Password      string
	PhoneNumber   string
	FullName      string
	Role          int32
	Province      string
	City          string
	ZipCode       int32
	AddressDetail string
	PhotoURL      string
	Description   string
}
