//as we saw we can declare a list of variables using var in last main.go
//we can also initialize during declaration and omit type declaration

package main

import (
	"fmt"
)

var i, j, k = 1, "Hello", true

//locally in a function we can use := instead of var to declare and initialize a variable line 13

func main() {
	x, y, z := 11, "habibi", true
	fmt.Println("Demonstrantion of short declaration :", x, y, z)
	fmt.Println(i, j, k)

	rr := 33
	var f uint = uint(rr) //you have to use conversion function to use data of one data type in other,,,,,

	fmt.Println(f)

	//Derclaration and usage of constant
	const name = "Shreyas" // automatically gets the type and also we cannot declare it using := and once declared cant be changed
	fmt.Println("My name is ", name)

	// Call the loop function
	loop()

	//if statements run com, refer ifstatements.go
	state()

	//Switch case 

	swit()

	//Practice 
	fmt.Println("Practice of Go")
	pract()

	//Pointers
	point()

}
