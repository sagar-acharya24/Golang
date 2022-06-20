package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to the time study of golang")

	presentTime := time.Now()

	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:06:02 Monday"))

	fmt.Println("Created Time")

	createdTime := time.Date(2021, time.April, 24,16,12,30,0,time.UTC)

	fmt.Println(createdTime)

	fmt.Println(createdTime.Format("01-02-2006 14:12:30 Monday"))
}
