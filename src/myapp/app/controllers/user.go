package controllers

import (
    "github.com/revel/revel"
)

type User struct {
    *revel.Controller
}

func (c User) Index(id string) revel.Result {
    return c.Render(id)
}
