package controllers

import (
	"encoding/json"
	"net/http"
	"paulocenteno/microsrvcs/services"
	"strconv"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.ParseInt(r.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("user_id must be an integer number"))
		return
	}
	user, err := services.GetUser(userId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	jsonValue, _ := json.Marshal(user)
	w.Write(jsonValue)
}
