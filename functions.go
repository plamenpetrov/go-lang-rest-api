package main

import (
	"fmt"
)

func mainFunction() {
	port := 3000
	//isStarted := startWebServer(port, 2)
	//fmt.Println(isStarted)

	p, err := startWebServer(port, 1)
	fmt.Println(p, err)

	_, err2 := startWebServer(port, 2)
	fmt.Println(err2)
}

func startWebServer(port int, numberOfRetries int) (int, error) {
	fmt.Println("Start server on port", port)
	fmt.Println("Number of retries", numberOfRetries)

	return port, nil
}