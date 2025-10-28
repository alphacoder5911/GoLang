package main

import (
	"fmt"
	"math"

	"rsc.io/quote"
)

func add(x int, y int) int {
	return x + y
}

func sub(x, y int) int { // if two or more params have same type, you can use this syntax
	return x - y
}

// A function can return multiple values and suppose u have given to swap two strings you can use this .

func swap(x, y string) (string, string) {

	return y, x
}

// Naked return , return values are already affixed at begining of the func when we declare return type , in end if we only use return it auto return x,y , not a recommended way to return for long functions ,
// can affect readability and maintainability.
func nakedReturn(sum int) (x, y int) {
	x = sum + 50
	y = 30 + sum
	return
}

var cat, dog, python bool // var declares a list of variables which can be used later , if declared globally like this can be used in package level and in local function if declared there

func main() {

	var i, j, k int

	j = 22
	k = 33

	fmt.Println(j, k)

	fmt.Printf("Hello, shreyas. ")
	fmt.Println(quote.Go())
	fmt.Println(math.Sqrt(35))

	//fUNCTION CALL
	fmt.Println(add(10, 20))

	//Swapping
	a, b := swap("Hello", "World")
	fmt.Println(a, b)

	//naked return
	fmt.Println(nakedReturn(10))

	//var usage

	fmt.Println(i, cat, dog, python)

}
