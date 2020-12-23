package main

func mainLoops()  {
	var i int

	for i < 5 {
		println(i)
		i++

		if (i == 3) {
			break
		}
	}

	for ; i < 5; i++ {
		println(i)
	}

	//INFINIT LOOP (condition added)
	for ; ; {
		if i == 5 {
			break
		}

		println(i)
		i++
	}

	//LOOP Collections
	println("================================================================================")
	slice := []int{1, 2, 3}

	for i, v := range slice {
		println(i, v)
	}

	println("================================================================================")
	wellKnownPorts := map[string]int{"http": 80, "https": 443}
	for protocol, value := range wellKnownPorts {
		println(protocol, value)
	}

}