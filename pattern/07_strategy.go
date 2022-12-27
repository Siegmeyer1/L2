package pattern

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

type IExecutor interface {
	Execute()
}

type Executor struct {
	strategy IExecutor
}

func (e *Executor) Execute() {
	e.strategy.Execute()
}

type strategy1 struct{}

func (s *strategy1) Execute() {
	fmt.Println("Executed strat #1")
}

type strategy2 struct{}

func (s *strategy2) Execute() {
	fmt.Println("Executed strat #2")
}

func example_strategy() {
	executor := &Executor{&strategy1{}}
	executor.Execute()
	executor.strategy = &strategy2{}
	executor.Execute()
}
