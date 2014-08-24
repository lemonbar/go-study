package lemon

import "fmt"

type Struct1 struct {
	Int1 int
}

func ChangeValue(value *Struct1, newValue int) {
	value.Int1 = newValue
}

func NotChangeValue(value Struct1, newValue int) {
	value2 := &value
	value2.Int1 = newValue
}

func Min(ints ...int) int {
	if len(ints) == 0 {
		return -1
	}
	min := ints[0]
	for _, x := range ints {
		if x < min {
			min = x
		}
	}
	return min
}

func PrintType(variables ...interface{}) {
	for _, v := range variables {
		switch v.(type) {
		case int:
			fmt.Println("type is int %d", v)
		default:
			fmt.Println("other type %v", v)
		}
	}
}

func DeferOrder() {
	for i := 0; i < 4; i++ {
		defer fmt.Println("index value is %d", i)
	}
}

func DeferA() {
	//defer fmt.Println("defer in DeferA")
	fmt.Println("in DeferA")
	deferB()
	defer fmt.Println("defer in DeferA")
}

func deferB() {
	defer fmt.Println("defer in DeferB")
	fmt.Println("in DeferB")
}
