package models

import (
	"fmt"
	"github.com/revel/revel"
	"regexp"
)

// Simple User Class
type User struct {
    UserId int
    Name string
}

func (u *User) String() string {
    return fmt.Sprintf("User(%s)", u.Name)
}
