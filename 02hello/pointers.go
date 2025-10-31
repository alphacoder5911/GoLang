package main 

import "fmt"

func point(){

    // similar to C but Go does not support pointer arithmetic
    i := 33
    p := &i // p stores the address of i
    fmt.Println("address:", p, "value:", *p)

	fmt.Println("Now changing value through pointer")
	*p=555
	fmt.Println("address:",p,"value",*p)



	j:=125

	z:=&j
	fmt.Println("address:",z,"value",*z)

    *z=*z/5
    fmt.Println("address:", z, "value", *z, "\n")
	
	//Practice 
	fmt.Println("Practice \n")
	a:=10
	b:=5
	c:=5

	d:=&a
	e:=&b
	f:=&c 

	fmt.Println("address:",d,"value",*d)
	fmt.Println("address:",e,"value",*e)
	fmt.Println("address:",f,"value",*f)

	fmt.Println("Math operations on em \n")

	*d=(*d-*e)/ *f

	fmt.Println("address:",d,"value",*d)




}