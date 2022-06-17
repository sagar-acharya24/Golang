package main

import (
	"fmt"
	"time"
)

type friend struct {
	course string
}

func friendship(f chan string) {

	var yashwanth friend = friend{"golang"}

	var rajkumar friend = friend{"golang"}

	fmt.Println(<-f, yashwanth, rajkumar)
}

var loan chan info = make(chan info, 10) //buffered channel and declared globally

func main() {
	fmt.Println("human relationship")

	jvt := make(chan string)

	func() {
		fmt.Println("relationship")
	}()

	go friendship(jvt)

	fmt.Println("write into channel")

	jvt <- "come for course"

	go house()

	func() {
		fmt.Println("bank")
		time.Sleep(2)
		fmt.Println("loan apprval", <-loan)
	}()

	fmt.Println("main function execution")
}

type info struct {
	owner string
	size  int
	floor int
}

func bank() {

}

func house() {
	var sagar info = info{"rama", 1200, 3}

	loan <- sagar
}
