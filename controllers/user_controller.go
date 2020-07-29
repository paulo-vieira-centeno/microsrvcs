package controllers

import (
	"encoding/json"
	"net/http"
	"paulocenteno/microsrvcs/services"
	"paulocenteno/microsrvcs/utils"
	"strconv"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.ParseInt(r.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		appError := &utils.AppError{
			Msg:    "user_id must be an integer number",
			Status: http.StatusBadRequest,
			Code:   "bad_request",
		}
		jsonValue, _ := json.Marshal(appError)
		w.WriteHeader(appError.Status)
		w.Write(jsonValue)
		return
	}
	user, appError := services.GetUser(userId)
	if appError != nil {
		jsonValue, _ := json.Marshal(appError)
		w.WriteHeader(appError.Status)
		w.Write([]byte(jsonValue))
		return
	}
	jsonValue, _ := json.Marshal(user)
	w.Write(jsonValue)
}
