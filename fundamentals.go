package main

import "fmt"

const (
	firstConst = 1
	secondConst = 2
)
func mainFundamentals() {
	//PRIMITIVES
	fmt.Println("Hello, Paco")

	var i int
	fmt.Println(i)

	var f float32 = 3.14
	fmt.Println(f)

	firstName := "Plamen"
	fmt.Println(firstName)

	b := true
	fmt.Println(b)

	c := complex(3, 4)
	fmt.Println(c)

	//new assignment
	r, im := real(c), imag(c)
	fmt.Println(r, im)

	//END PRIMITIVES

	//POINTERS
	var name *string = new(string)

	*name = "Plamen"

	//access pointer memory address => 0xc00003e200
	fmt.Println(name)

	//access pointer value => Plamen
	fmt.Println(*name)

	firstVariableName := "My name"
	fmt.Println(firstVariableName)

	pointerToVariableName := &firstVariableName
	fmt.Println(pointerToVariableName, *pointerToVariableName)

	firstVariableName = "Paco"
	fmt.Println(pointerToVariableName, *pointerToVariableName)		//memory address is the same but value is different

	//END POINTERS

	//CONSTANTS
	const consant int = 3

	fmt.Println(consant + 3)
	fmt.Println(float32(consant) + 1.2)

	fmt.Println(firstConst, secondConst)

	//END CONSTANTS
}
