package domain

import (
	"net/http"
	"paulocenteno/microsrvcs/utils"
	error_utils "paulocenteno/microsrvcs/utils"
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

func GetUser(userId int64) (*User, *error_utils.AppError) {
	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, &utils.AppError{
		Msg:    "User not found",
		Status: http.StatusNotFound,
		Code:   "not_found",
	}
}
