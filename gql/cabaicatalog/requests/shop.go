package requests

type ListShop struct {
	Page  int32
	Limit int32
}

type GetShop struct {
	ID       int32
	SlugName string
}
