package main

import "fmt"

func pract() {

	sum := 0

	for sum < 100 {
		sum += 100
	}
	fmt.Println("The sum is ",sum)

	//Multiplication

	for i:=1;i<=10;i++{
		fmt.Println("5 x",i,"=",5*i)
	}

	//ODD or Even

	for i:=0 ;i<30;i++ {
		if(i%2==0){
			fmt.Println(i)
		}
	}

}