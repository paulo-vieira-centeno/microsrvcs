package app

import (
	"net/http"
	"paulocenteno/microsrvcs/controllers"
)

func RunApp() {
	http.HandleFunc("/users", controllers.GetUser)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
