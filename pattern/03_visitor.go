package pattern

import "fmt"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

type someComplexStruct struct {
	a int
	b string
}

type IVisitor interface {
	visit(*someComplexStruct)
}

type VisitorA struct{}
type VisitorB struct{}

func (v VisitorA) visit(scs *someComplexStruct) {
	scs.a += 1
	fmt.Println(scs.a)
}

func (v VisitorB) visit(scs *someComplexStruct) {
	scs.b += "1"
	fmt.Println(scs.b)
}

func (scs *someComplexStruct) Accept(v IVisitor) {
	v.visit(scs)
}

func example_visitor() {
	str := someComplexStruct{a: 1, b: "hello"}
	str.Accept(VisitorA{})
	str.Accept(VisitorB{})
}
