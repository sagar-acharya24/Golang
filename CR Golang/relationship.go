
package main

import (
	"fmt"
	"time"
)

type friend struct{
	course string
}

func friendship(f chan string){
	var yash friend=friend{"golang"}
	var raj friend=friend{"golang"}
	fmt.Println(<-f,yash,raj)
}

func main(){
	fmt.Println("human relation")

	jvt := make(chan string)

	func (){
		fmt.Println("relationship")

	}()

	go friendship(jvt)

	fmt.Println("write into channel")

	jvt <- "come for course"

	time.Sleep(1)
}
