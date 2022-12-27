package pattern

import "fmt"

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

type complexStruct struct {
	complexInt    int
	complexString string
}

func (cs *complexStruct) setComplexInt(input int) {
	cs.complexInt = input
}

func (cs *complexStruct) setComplexString(input string) {
	cs.complexString = input
}

type anotherComplexStruct struct{}

func (acs *anotherComplexStruct) doComplexStuff(cs *complexStruct) {
	fmt.Println(string(rune(cs.complexInt)) + "->" + cs.complexString)
}

type EasyfyingFacade struct {
	cmplStruct      *complexStruct
	anthrCmplStruct *anotherComplexStruct
}

func NewEasyfyingFacade() *EasyfyingFacade {
	return &EasyfyingFacade{cmplStruct: &complexStruct{}, anthrCmplStruct: &anotherComplexStruct{}}
}

func (ef *EasyfyingFacade) doComplexStuffEasy(easyInt int, easyString string) {
	ef.cmplStruct.setComplexInt(easyInt)
	ef.cmplStruct.setComplexString(easyString)
	ef.anthrCmplStruct.doComplexStuff(ef.cmplStruct)
}
