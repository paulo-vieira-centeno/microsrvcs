package domain

import (
	"errors"
	"fmt"
)

var (
	users = map[int64]*User{
		123: &User{Id: 1,
			FName: "Zion",
			LName: "Phelps",
			Email: "zion@phelps.com",
		},
	}
)

func GetUser(userId int64) (*User, error) {
	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, errors.New(fmt.Sprintf("User %v not found!", userId))
}
