package main

import (
	"fmt"
	"github.com/pluralsight/webservice/models"
)

func mainService() { 			//must be renamet in order to work
	u := models.User{
		ID: 2,
		FirstName: "Plamen",
		LastName: "Petrov",
	}

	fmt.Println(u)
}