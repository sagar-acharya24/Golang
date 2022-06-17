package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	fmt.Println("before second")
	go func() {
		fmt.Println("second go routine")
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	fmt.Println("before first")
	go func() {
		fmt.Println("first go routine")
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	fmt.Println("before loop")
	for i := 0; i < 2; i++ {
		fmt.Println("before selection")
		select {
		case msg1 := <-c1:
			fmt.Println("mesaaged", msg1)
		case msg2 := <-c2:
			fmt.Println("messaged", msg2)
		}
	}
}
