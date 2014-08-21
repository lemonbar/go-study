package lemon

import "strings"

type Person struct {
	FirstName string
	LastName  string
}

type InnerS struct {
	In1 int
	In2 int
}

type Int int

type Outer struct {
	B      int
	C      float32
	Int    //anonymous field
	InnerS //anonymous field
}

func (p *Person) UpPerson() {
	p.FirstName = strings.ToUpper(p.FirstName)
	p.LastName = strings.ToUpper(p.LastName)
}

func (p *Person) String() string {
	return "FirstName is " + p.FirstName + "; LastName is " + p.LastName
}
