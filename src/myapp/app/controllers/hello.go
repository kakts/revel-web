package controllers

import (
	"github.com/revel/revel"
)

type HelloApp struct {
	*revel.Controller
}

func (c HelloApp) Index() revel.Result {
	greeting := "Aloha World"
	return c.Render(greeting)
}
