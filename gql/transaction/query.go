package transaction

const (
	Query = `
		# List Carts
		carts(params: ListCartInput!): [Cart]
		customerOrders(params: ListOrderInput!): [Order]
		shopOrders(params: ListOrderInput!): [Order]
	`
	Mutation = `
		createCart(params: CreateCartInput!): Cart
		updateCartQuantity(params: UpdateCartInput!): Cart
		deleteCart(cartID: Int!): Success
		
		checkout(params: CheckoutInput!): [Order]
		shipOrder(orderID: Int!): Order
		rejectOrder(orderID: Int!): Order
		fulfillOrder(orderID: Int!): Success
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
		input ListSaldoHistoryInput{
			page: Int = 1
			limit: Int = 10
		}
		input ListCartInput{
			page: Int = 1
			limit: Int = 10
		}
		input ListOrderInput{
			status: String = ""
			page: Int = 1
			limit: Int = 10
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
