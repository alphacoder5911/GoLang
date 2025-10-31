package main

import (
	"fmt"
	"runtime"
	"time"
)

func swit() {

	fmt.Println("This is switch case demonstration")

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("MacOs")

	case "linux":
		fmt.Println(("Linux"))

	default:
		fmt.Println("%s. \n", os)
	}

	fmt.Println("This is a switch without condition , i.e switch is always true ")
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning")

	case t.Hour() > 12 && t.Hour() < 16:
		fmt.Println("Good afternoon")

	default:
		fmt.Println("Good evening")
	}

	//switch parctice

	age := 20

	switch  {
	case age > 18:
		fmt.Println("Adult")
	default:
		fmt.Println("Minor")
	}

	mul:=2

	switch ans:=5 * mul;ans{
	case 20 : 
		fmt.Println(ans)
	case 10 :
		fmt.Println(ans)

	default :
		fmt.Println(ans)
	}

	//Defer : the will exec the defered function but will wait till other nearby non defered functions to return 

	defer fmt.Println("This Is Shreyas")
	defer fmt.Println("Class")
	fmt.Println("Hello")
	fmt.Println("----")

	//if there are multipe differ then go will exec them in lifo manner.

	//Deferred function calls are pushed onto stack , hence executing in lifo manner 

	fmt.Println("counting")
	for i:=0;i<10;i++{
		defer fmt.Println(i)
	}
	//The output of the above loop is an evidence to the statement.
	fmt.Println("Done ")


}
