/*package main

import ("fmt")

type marriage struct{
	dowry string
	gold string
	car string
	house string
	salary int
}

func main(){
	var ajay marriage = marriage{"5cr","1kg","30lakhs","duplex",70000}

	var aj marriage = marriage{"10cr","1.5kg","30lakhs","duplex",70000}

	fmt.Println(ajay,aj)
}*/

/*package main

import "fmt"

type happiness struct {
	selfhappy string
	enjoy     string
}

type richlife struct {
	happiness
	job   string
	money string
}

type familylife struct {
	mariage string
}

func main() {
	var mj richlife = richlife{happiness{"yes", "yes"}, "rs50000", "1cr"}
	var vb familylife = familylife{"yes"}
	fmt.Println(mj,vb)
}*/

package main

import "fmt"

type hurting interface{
	nickname string
	teasing  string
	respect  string
}


type school struct {
	nickname string
	teasing  string
	respect  string
}

func () nickname() {

}

func (h string) teasing(t string) {
	h.teasing = t
}

func () respect() {

}

type college struct {
	nickname string
	teasing  string
	respect  string
}

func () nickname() {

}

func () teasing() {

}

func () respect() {

}

type job struct {
	nickname string
	teasing  string
	respect  string
}

func () nickname() {

}

func () teasing() {

}

func () respect() {

}

func main() {
	var anil school
	anil.teasing("short fellow")
	fmt.Println(anil.teasing)
}


