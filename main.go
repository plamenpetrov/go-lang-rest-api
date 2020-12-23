package main

import (
	"net/http"
	"github.com/pluralsight/webservice/controllers"
	//"github.com/pluralsight/webservice/models"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}