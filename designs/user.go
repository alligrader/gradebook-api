package designs

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("user", func() {
	BasePath("/user")
	Description("Represents a user of the site")

	// Corresponds to creating a new user
	Action("create", func() {
		Description("Sign up for the first time")
		Routing(POST("/"))
		Payload(UserCreate)
		Response(OK, func() {
			Status(201)
			Media(UserMedia)
		})
		Response(InternalServerError)
		Response(Conflict)
	})

	// Login
	Action("login", func() {
		Description("Log in to a user account")
		Routing(POST("/login"))
		Payload(Login)
		Response(OK, UserMedia)
		Response(Unauthorized)
	})

	Action("update", func() {
		Description("Adjust your account settings")
		Routing(PUT("/:userID"))
		Payload(UserCreate)
		Response(OK, UserMedia)
		Response(InternalServerError)

		Params(func() {
			Param("userID", Integer, "The user's unique identifier")
		})
	})

	Action("read", func() {
		Description("Get the detials about this particular account")
		Routing(GET("/:userID"))
		Response(OK, UserMedia)
		Response(InternalServerError)

		Params(func() {
			Param("userID", Integer, "The user's unique identifier")
		})
	})
})

var UserMedia = MediaType("application/user.mt", func() {
	Reference(UserCreate)
	Attributes(func() {
		Attribute("id", Integer)
		Attribute("name")
		Attribute("email")
		Attribute("password")
		Attribute("token")
		Attribute("expiry")
		Attribute("ghUsername")
	})
	View("default", func() {
		Attribute("id", Integer)
		Attribute("name")
		Attribute("email")
	})

	View("Incoming", func() {
		Attribute("name")
		Attribute("email")
		Attribute("password")
	})
	View("OnLogin", func() {
		Attribute("id", Integer)
		Attribute("name")
		Attribute("email")
		Attribute("token")
		Attribute("expiry")
	})
})

var UserCreate = Type("UserCreate", func() {
	Description("payload used to sign up a user")
	Attribute("name", func() {
		MinLength(1)
	})
	Attribute("email", String, "User email", func() {
		Format("email")
		MinLength(1)
	})
	Attribute("password", func() {
		MinLength(2)
	})
})

var Login = Type("Login", func() {
	Description("represents the login creditentials which are sent to the server")
	Attribute("email", String, "User email", func() {
		Format("email")
		MinLength(1)
	})
	Attribute("password", func() {
		MinLength(2)
	})
})
