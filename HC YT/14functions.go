
package main

import "fmt"

func main() {
	fmt.Println("welcome to functions in golang")

	greeter()

	result := adder(3,5)

	fmt.Println("Result is : ",result)

	proRes,myMessage := proAdder(2,3,5,7,8,5)

	fmt.Println("Proreslt is :", proRes)

	fmt.Println("My message is : ", myMessage)
}

func greeter() {
	fmt.Println("Namasthe in golang")
}

func adder(valOne int, valTwo int)int {
	return valOne + valTwo	
}

func proAdder(values ...int)(int,string){
	total :=0
	
	for _,val := range values{
		total += val
	}
	return total, "Hi ProResult in golang"
}
