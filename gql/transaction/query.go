package transaction

const (
	Query = `
		# List Carts
		carts(token: String = ""): [Cart]
		customerOrders(token: String = ""): [Order]
		shopOrders(token: String = ""): [Order]
	`
	Mutation = `
		createCart(params: CreateCartInput!): Cart
		updateCartQuantity(params: UpdateCartInput!): Cart
		deleteCart(cartID: Int!): Success
		
		checkout(params: CheckoutInput!): [Order]
		shipOrder(orderID: Int!): Order
`
	Types = `
		type Cart{
			id: ID!
			product:  Product
			userID: Int!
			AmountKG: Float!
		}
		type Order{
			id: ID!
			customerID: ID!
			shopID: ID!
			totalPrice: String!
			products: [Product]
			status: String!
			payment: Payment
		}
		type Payment{
			id: ID!
			amount: String!
			status: String!
			method: String!
			createdAt: String!
			updatedAt: String!
		}
		input CreateCartInput{
			productID: Int!
			quantityKG: Float!
		}
		input UpdateCartInput{
			cartID: Int!
			newQuantityKG: Float!
		}
		input CheckoutInput{
			cartIDs: [String!]!
			paymentAmount: String!
		}
`
)
