package pattern

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

type IStage interface {
	Run(int)
	SetNext(IStage)
}

type Qube struct {
	next IStage
}

func (q *Qube) SetNext(next IStage) {
	q.next = next
}

func (q *Qube) Run(val int) {
	val = val ^ 3
	q.next.Run(val)
}

type MinusTwo struct {
	next IStage
}

func (m *MinusTwo) SetNext(next IStage) {
	m.next = next
}

func (m *MinusTwo) Run(val int) {
	val -= 2
	fmt.Println(val)
}

func example_chain() {
	Q := Qube{}
	Q2 := Qube{}
	M := MinusTwo{}
	Q2.SetNext(&M)
	Q.SetNext(&Q2)
	Q.Run(2)

}
