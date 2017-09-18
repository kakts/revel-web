package models

import (
    "log"
	"github.com/revel/revel"
	"regexp"
)

// Simple User Class
type User struct {
    UserId      string
    Name        string
    Password      string
    HashedPassword     []byte
}

func (u *User) String() string {
    return "test"
}

var userRegex = regexp.MustCompile("^\\w*$")

func (user *User) Validate(v *revel.Validation) {
    log.Printf("User(%s,%s,%s)", user.UserId, user.Name, user.Password)
    v.Check(user.UserId,
        revel.Required{},
        revel.MinSize{4},
        revel.MaxSize{20},
    )
    v.Check(user.Name,
        revel.Required{},
        revel.MaxSize{15},
        revel.MinSize{4},
    )

    ValidatePassword(v, user.Password).Key("user.Password").Message("Add user.Password")
}

func ValidatePassword(v *revel.Validation, password string) *revel.ValidationResult {
    return v.Check(password,
        revel.Required{},
        revel.MaxSize{15},
        revel.MinSize{5},
    )
}
