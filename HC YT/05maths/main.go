
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Hello maths in golang")

	var numone int = 4
	var numtwo float64 = 6

	fmt.Println("Sum is : ", numone+int(numtwo))

	// rand number from math/rand package
	rand.Seed(time.Now().UnixNano())

	fmt.Println(rand.Intn(5) + 1)

	// random number from rand/crypto package
	// more reliable because of crypto

	//randNum, err := rand.Int(rand.Reader,big.NewInt(5))

	
	//fmt.Println(randNum)
}
