package main

import (
	"net/http"

	"github.com/komoda102/learningGo_usersCRUDWebAPI/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
