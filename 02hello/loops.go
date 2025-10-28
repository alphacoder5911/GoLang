package main

import (
	"fmt"
)

func loop() {

	sum := 0

	for i := 0; i < 20; i++ {
		sum += 1
	}

	fmt.Println("Sum is ", sum)

	for i := 0; i < 10; i++ {

		for j := 0; j < 20; j++ {
			fmt.Print("*")
		}
		fmt.Print("\n")

	}

	// We can also omit init and post statements
	suwm := 1
	for suwm < 100 {
		suwm += suwm
	}

	fmt.Println(suwm)

	//SO as for while is concerned , for loop is itself while in go , c's while is spelled for in go 😂
	small := 0
	large := 10
	for small < large {
		small += 1

	}

	fmt.Println(small)

	//You can also create a never ending for loog which god knows why we need  --> for { } --> infinite shit

	//Next if statements ....

}
