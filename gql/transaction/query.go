package transaction

const (
	Query = `
		# List Carts
		carts(token: String = ""): [Cart]
	`
	Mutation = `
		createCart(params: CreateCartInput!): Cart
		updateCartQuantity(params: UpdateCartInput!): Cart
`
	Types = `
		type Cart{
			id: ID!
			product:  Product
			userID: Int!
			AmountKG: Float!
		}
		input CreateCartInput{
			productID: Int!
			quantityKG: Float!
		}
		input UpdateCartInput{
			cartID: Int!
			newQuantityKG: Float!
		}
`
)
