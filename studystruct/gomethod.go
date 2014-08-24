package lemon

import "fmt"

type MethodStruct01 struct {
	Int1 int
}

func (ms *MethodStruct01) Method1() {
	fmt.Println("this is method01")
}
