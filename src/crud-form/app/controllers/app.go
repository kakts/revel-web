package controllers

import (
	"github.com/revel/revel"

    "log"
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

func init() {

}
