package controllers

import (
	"github.com/revel/revel"

    "log"

    "crud-form/app/models"
    "crud-form/app/routes"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
    // Get the query parameter.
    test := c.Params.Query.Get("test")
    log.Print(test)
	return c.Render()
}

func (c App) Test() revel.Result {
    // Get the query parameter.
    test := c.Params.Form.Get("ui")
    c.Validation.MinSize(test, 5)
    log.Print(test)
	return c.Render()
}

func (c App) Register() revel.Result {
    return c.Render()
}

func (c App) SaveUser(user models.User, verifyPassword string) revel.Result {
    c.Validation.Required(verifyPassword)
    c.Validation.Required(verifyPassword == user.Password).Message("Password does not match")
    user.Validate(c.Validation)

    // Error時の処理
    if c.Validation.HasErrors() {
        c.Validation.Keep()
        c.FlashParams()
        return c.Render(routes.App.Index())
    }


    c.Session["user"] = user.Username
    c.Flash.Success("Welcome, " + user.Name)

    // TODO routerがおかしい
    return c.Render(routes.App.Index())
}
