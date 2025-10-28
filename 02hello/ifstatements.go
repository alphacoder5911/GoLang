package main

import "fmt"



func divide(x,y int) int{
	if y!=0{
		return x/y
	}else{
		fmt.Println("Cannot divide by 0")
	}

	return 0
}

// Add your if statement examples here
func state(){
	//same like for loops no parenthesis required , but braces are required 

	x:=10
	if x >5{
		fmt.Println("Shreyas says hello")
	}

	left:=10
	right:=5
	//lets take a functions example 
	fmt.Println(` ${left} divided by ${right}`,divide(left,right))

	// if with a short statement i.e a:=11 is called a short statement avoiding explicit declarations 
	var lim int =20
	if v:= 5*6; v < lim {
		fmt.Println(v)
	}
}