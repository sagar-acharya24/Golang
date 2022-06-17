package main

import "fmt"

func main() {
	fmt.Println("Welcome to class of pointers")

	//var ptr *int

	//fmt.Println("Value of pointer is : ", ptr)

	mynum := 24

	var ptr = &mynum

	fmt.Println("Value of pointer memory address: ", ptr)

	fmt.Println("Value of actual pointer: ", *ptr)

	*ptr = *ptr + 2

	fmt.Println("New Value is : ", mynum)
}
