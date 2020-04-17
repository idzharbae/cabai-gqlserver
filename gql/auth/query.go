package auth

const (
	Query = `
		login(params: LoginInput!): Token
		refreshToken(params: RefreshTokenInput!): Token
		getUserInfo(token: String): User
	`
	Mutation = `
		register(params: RegisterInput!): User
		editProfile(params: EditProfileInput!): User
`
	Types = `
		type Success{
			success: Boolean!
		}
		type Token{
			token: String!
		}
		type User{
			id:      ID!
			name: String!
			userName: String!
			email: String!
			phone: String!
			type: Int!
			photoURL: String!
			city: String!
			province: String!
			zipCode: Int!
			addressDetail: String!
			description: String!
			createdAt: String!
			updatedAt: String!
		}
		input LoginInput{
			userNameOrEmail: String!
			password: String!
		}
		input RefreshTokenInput{
			token: String!
		}
		input RegisterInput{
			userName: String!
			email: String!
			password: String!
			phoneNumber: String!
			fullName: String!
			role: Int!
			city: String = ""
			province: String = ""
			addressDetail: String = ""
			zipCode: Int = 0
			description: String = ""
		}
		input EditProfileInput {
			password: String!
			newPassword: String!
			phoneNumber: String!
			fullName: String!
			city: String = ""
			province: String = ""
			addressDetail: String = ""
			zipCode: Int = 0
			description: String = ""
			photo: Upload
		}
`
)
