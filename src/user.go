package main

import (
	"log"

	"github.com/alligrader/gradebook-api/src/app"
	_ "github.com/go-sql-driver/mysql"
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"

	"github.com/alligrader/gradebook-api/src/models"
)

// UserController implements the user resource.
type UserController struct {
	*goa.Controller
}

// NewUserController creates a user controller.
func NewUserController(service *goa.Service) *UserController {
	return &UserController{Controller: service.NewController("UserController")}
}

// Create runs the create action.
func (c *UserController) Create(ctx *app.CreateUserContext) error {
	// UserController_Create: start_implement

	// Open the connection to the DB
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/alligrader")
	db.AutoMigrate(&models.User{})
	if err != nil {
		log.Println(err)
		return ctx.InternalServerError()
	}
	userDB := models.NewUserDB(db)

	// Fetch the user out of the request context
	u := models.UserFromuser(ctx.Payload)

	// Add that user to the database
	userDB.Add(ctx, u)

	// UserController_Create: end_implement
	return ctx.OK(u.UserToUserMt())
}