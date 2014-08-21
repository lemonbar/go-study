// main
package main

import (
	"fmt"

	lemon "gostudy/studystruct"
)

func main() {
	showStructInitialUsage()
	showStructAnonyFieldUsage()
}

func showStructAnonyFieldUsage() {
	outer := new(lemon.Outer)
	outer.In1 = 3 //outerS.in1 is innerS.in1
	outer.In2 = 4 //outerS.in2 is innerS.in2
	outer.Int = 5 //int is also the anonymous filed name

	//outer2 := lemon.Outer{5, 6.7, 20, lemon.InnerS{3, 4}} //outerS can also be initialized in this format
}

func showStructInitialUsage() {
	//1 - struct as a value type:
	var pers1 lemon.Person
	pers1.FirstName = "Lemon"
	pers1.LastName = "Bar"
	fmt.Println("before to upper ...")
	fmt.Println(&pers1)
	(&pers1).UpPerson()
	fmt.Println("after to upper ...")
	fmt.Println(&pers1)

	//2 - struct as a pointer:
	pers2 := new(lemon.Person)
	pers2.FirstName = "Lemon"
	pers2.LastName = "Bar"
	//(*pers2).lastName = "Bar"		//this is also valid
	pers2.UpPerson()

	//3 - struct as a literal:
	pers3 := &lemon.Person{"Lemon", "Bar"}
	//pers3 := &Person{firstName:"Lemon",lastName:"Bar"}	//this is also valid
	pers3.UpPerson()
}
