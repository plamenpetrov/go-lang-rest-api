package main

import (
	"fmt"
)

func mainCollections() {
	var arr [3]int

	arr[0] = 1
	arr[1] = 2
	arr[2] = 3

	fmt.Println(arr)

	arr2 := [3]int{1,2,3}
	fmt.Println(arr2)

	//SLICE
	fmt.Println("SLICE =========================================================================")
	slice := arr2[:]

	arr2[0] = 332
	slice[2] = 432

	fmt.Println(arr2, slice)
	
	//IMPORTANT: This not work with just array, ONLY with slice
	slice = append(slice, 5, 6, 7)
	fmt.Println(slice)

	s2 := slice[1:]			// start from first element to the end of array
	s3 := slice[:2]			// start from the begining and get 2 elements
	s4 := slice[1:2]		// start from the first element and get 2 elements

	fmt.Println(s2, s3, s4)

	//END SLICE

	//MAP
	fmt.Println("MAP =========================================================================")

	m := map[string]int{"firstMapElement": 1, "secondMapElement": 2} 		//same as json string
	fmt.Println(m["firstMapElement"])

	delete(m, "FirstMapElement")
	fmt.Println(m)
	//END MAP

	//STRUCT
	fmt.Println("STRUCT =========================================================================")

	type user struct {
		ID int
		FirstName string
		LastName string
	}

	var u user
	u.ID = 1
	u.FirstName = "Plamen"
	u.LastName = "Petrov"

	fmt.Println(u)

	u2 := user{ ID: 1,					//space after { is VERY IMPORTANT with multiline sintax
		FirstName: "Paco",
		LastName: "Paco",			//comma is VERY IMPORTANT HERE
	}

	fmt.Println(u2)
	//END STRUCT
}
