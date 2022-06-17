package main

import (
	"fmt"
	//"time"
	"sync"
)

func dreams(k chan string, w *sync.WaitGroup) {

	fmt.Println("dreams may execute")

	k <- "rich developer"

	fmt.Println(" writing dream") //cannotadd and read together that is the mistake example fmt.Println("dream",<-k )

	//time.Sleep(3)

	fmt.Println("dreams after sleep")

	w.Done()

}

var kn map[string][]string = make(map[string][]string)

func knowledge(k1 chan string, w *sync.WaitGroup) {

	fmt.Println("knowledge")

	//var know chan string

	//know <-k1

	switch <-k1 {

	case "rich developer":
		kn["6 june"] = []string{"go routine", "channel"}
	case "doctor":
		kn["6 june"] = []string{"patient", "reading book", "pratical"}
	}

	fmt.Println("map data structure", kn)

	w.Done()
}

func money(w *sync.WaitGroup) {

	fmt.Println("money")

	w.Done()

}

func health(w *sync.WaitGroup) {

	fmt.Println("health")

	w.Done()

}

func main() {

	fmt.Println("hello world")

	var wg sync.WaitGroup

	var know chan string = make(chan string)

	for i := 0; i < 2; i++ {

		wg.Add(1)

		go dreams(know, &wg)

		fmt.Println("for loop entry")

		//fmt.Scanln(&know) //console is a resource it blocks
		wg.Add(1)

		go knowledge(know, &wg)

		wg.Add(1)

		go money(&wg)

		wg.Add(1)

		go health(&wg)

		fmt.Println("now all goroutine may execute")

		wg.Wait()

		fmt.Println("for loop exit")
	}
}
