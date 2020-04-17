package requests

import upload "github.com/smithaitufe/go-graphql-upload"

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
	Description   string
}

type EditProfile struct {
	Password      string
	NewPassword   string
	PhoneNumber   string
	FullName      string
	Province      string
	City          string
	ZipCode       int32
	AddressDetail string
	Description   string
	Photo         *upload.GraphQLUpload
}
