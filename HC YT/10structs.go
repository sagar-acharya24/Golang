package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")

	sagar := User{"Sagar", "sagar@go.dev", true, 18}

	fmt.Println(sagar)

	fmt.Printf("Sagar details are: %+v\n", sagar)

	fmt.Printf("Name is %v and Email is %v", sagar.Name, sagar.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
