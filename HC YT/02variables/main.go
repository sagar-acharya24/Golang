package main

import "fmt"

const Loginvar string = "ghabbiji"

func main() {
	var user string = "sagar"
	fmt.Println(user)
	fmt.Printf("Variable is of type: %T \n",user)

	var isloggedin bool = false
	fmt.Println(isloggedin)
	fmt.Printf("Variable is of type: %T \n",isloggedin)

	var smallval int = 255
	fmt.Println(smallval)
	fmt.Printf("Variable is of type: %T \n",smallval)

	var smallfloat float64 = 255.44566674337
	fmt.Println(smallfloat)
	fmt.Printf("Variable is of type: %T \n",smallfloat)

	//default values n some alias
	var anval int
	fmt.Println(anval)
	fmt.Printf("Variable is of type: %T \n",anval)

	var web = "golang.org"
	fmt.Println(web)

	noofuser := 300000
	fmt.Println(noofuser)

	fmt.Println(Loginvar)
	fmt.Printf("Variable is of type: %T \n",Loginvar)
}
