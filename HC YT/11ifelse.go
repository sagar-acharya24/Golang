package main

import "fmt"

func main() {
	fmt.Println("If else in golang")

	loginCount := 24

	var result string

	if loginCount < 10 {
		result = "login count less than 10"
	} else if loginCount > 10 {
		result = "login more than 10"
	} else {
		result = "login exactly at 10"
	}
	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Num is even")
	} else {
		fmt.Println("Num is odd")
	}

	if num := 4; num <10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is more than 10")
	}
}
