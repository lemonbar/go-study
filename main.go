// main
package main

import (
	"fmt"
	"strconv"

	lemon "gostudy/studystruct"
)

func main() {
	//showStructInitialUsage()
	//showStructAnonyFieldUsage()
	//showFunctionTransferByValueAndRefer()
	//showFunctionMultiInputParameters()
	//showFunctionMultiInterfaceParameters()
	//showDeferUsage()
	//showMethodAliasNotContainsMethod()
	showInterfaceUsage()
}

func showInterfaceUsage() {
	sq1 := new(lemon.Square)
	sq1.Side = 5
	var areaIntf lemon.Shaper
	areaIntf = sq1

	//判断接口的类型
	if v, ok := areaIntf.(*lemon.Square); ok {
		fmt.Printf("The square has area: %f\n", v.Area())
	}

	//判断接口的类型
	switch areaIntf.(type) {
	case *lemon.Square:
		fmt.Printf("The square area is %f\n", areaIntf.Area())
	default:
		fmt.Printf("The default type")
	}

	//判断变量实现了接口
	if _, ok := areaIntf.(lemon.Shaper); ok {
		fmt.Printf("The square area is %f\n", sq1.Area())
	}
}

type Test lemon.MethodStruct01

func (test *Test) tm() {
	fmt.Println("new method")
}

func showMethodAliasNotContainsMethod() {
	var MS1 lemon.MethodStruct01
	MS1.Method1()
	var MS2 Test
	MS2.tm()
}

func showDeferUsage() {
	lemon.DeferOrder()
	//lemon.DeferA()
}

func showFunctionMultiInterfaceParameters() {
	lemon.PrintType(5, "aaaa")
	var2 := []interface{}{6, 7, 9, "bbb", "ccc"}
	lemon.PrintType(var2...)
}

func showFunctionMultiInputParameters() {
	fmt.Println(lemon.Min(4, 6, 9, 10, 3, 2))
	arr := []int{9, 5, 3, 10, 20, 2}
	fmt.Println(lemon.Min(arr...))
}

func showFunctionTransferByValueAndRefer() {
	changeValue := lemon.Struct1{Int1: 1}
	fmt.Println("before change value, Int1 value is " + strconv.Itoa(changeValue.Int1))
	lemon.ChangeValue(&changeValue, 2)
	fmt.Println("after change value, Int1 value is " + strconv.Itoa(changeValue.Int1))
	fmt.Println("before not change value, Int1 value is " + strconv.Itoa(changeValue.Int1))
	lemon.NotChangeValue(changeValue, 4)
	fmt.Println("after not change value, Int1 value is " + strconv.Itoa(changeValue.Int1))
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
