package controllers

import (
	"github.com/revel/revel"

    "log"

    "crud-form/app/models"
    "crud-form/app/routes"
    "crud-form/app/services"
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
        log.Fatal("Validation Error has occured.")
        for _, data := range c.Validation.Errors {
            log.Print(data.Message)
        }
        c.Validation.Keep()
        c.FlashParams()
        return c.Redirect(routes.App.Index())
    }
    results, errCode := services.FindOne(user.UserId, "testRevel", "User")
    if errCode != "" {
        panic(errCode)
    }
    log.Print(len(results))
    if len(results) > 0 {
        log.Fatal("UserId already exists")
        c.Flash.Error("UserId already exists.")
        return c.Redirect(routes.App.Index())
    }

    // TODO hashedpassword
    services.InsertEntity("testRevel", "User", user)

    c.Session["user"] = user.Name
    c.Flash.Success("Welcome, " + user.Name)

    // TODO routerがおかしい
    return c.Redirect(routes.App.Index())
}
/**
func (c App) getUser(userId string) *models.User {
    session, err := mgo.Dial("mongodb://192.168.33.10")
    if err != nil {
        panic(err)
    }
    defer session.Close()

    session.SetMode(mgo.Monotonic, true)

    return c.Render(routes.App.Index())
}

func (c App) Login (username string, password string, remember bool) revel.Result {
    user := c.getUser(username)
    if user != nil {
        err := bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(password))
        if err == nil {
            c.Session["user"] = username
            if remember {
                c.Session.SetDefaultExpiration()
            } else {
                c.Session.SetNoExpiration()
            }
            c.Flash.Success("Welcome, " + username)
            return c.Redirect(routes.App.Index())
        }
    }
}

func (c App) Logout() revel.Result {
    for k := range c.Session {
        delete(d.Session, k)
    }
    return c.Redirect(routes.App.Index())
}
*/
