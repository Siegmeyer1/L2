package pattern

import "fmt"

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

type LightBulb struct {
	isOn bool
}

func (lb *LightBulb) doStuffOn() {
	fmt.Println("Emits light")
}

func (ln *LightBulb) doStuffOff() {
	fmt.Println("Emits darkness")
}

func (lb *LightBulb) doStuff() {
	if lb.isOn {
		lb.doStuffOn()
	} else {
		lb.doStuffOff()
	}
}

func example_state() {
	lb := LightBulb{isOn: false}
	lb.doStuff()
	lb.isOn = true
	lb.doStuff()
}
