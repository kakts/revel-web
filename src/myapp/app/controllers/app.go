package controllers

import (
	"log"

	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Hello(myName string) revel.Result {
	// Add form validation
	c.Validation.Required(myName).Message("Your name is required!")
	c.Validation.MinSize(myName, 3).Message("Your name is not long enough!")

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Index)
	}

	return c.Render(myName)
}

// GET:/login
// TODO
func (c App) Login(id string, pass string) revel.Result {
	log.Println("test")
	if c.Validation.HasErrors() {
		log.Println("error")
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Index)
	}

	// success redirect to mypage
	// set cookie
	// TODO
	return c.Render()
}

// POST:/login
// TODO
func (c App) DoLogin(id string, pass string) revel.Result {
	log.Println("POST:/login")
	// validation for id
	c.Validation.Required(id).Message("id is required")
	c.Validation.MinSize(id, 8).Message("Id should be over 8 characters")

	// validation for id
	c.Validation.Required(pass).Message("pass is required")
	c.Validation.MinSize(pass, 8).Message("pass should be over 8 characters")
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Login)
	}

	// success redirect to mypage
	// set cookie
	// TODO
	return c.Render()
}
