package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")

	sagar := User{"Sagar", "sagar@go.dev", true, 18}

	fmt.Println(sagar)

	fmt.Printf("Sagar details are: %+v\n", sagar)

	fmt.Printf("Name is %v and Email is %v\n", sagar.Name, sagar.Email)

	sagar.getStatus()

	sagar.NewMail()

	fmt.Printf("Name is %v and Email is %v\n", sagar.Name, sagar.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) getStatus(){
	fmt.Println("Is user active:", u.Status)
}

func (u User) NewMail(){
	u.Email = "test@go.dev"
	fmt.Println("New Email is : ",u.Email) //This actually give the copy of this . So thats why pointer is used to give address or reference
}