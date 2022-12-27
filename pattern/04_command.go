package pattern

import "fmt"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

type ICmd interface {
	run()
}

type cmd struct {
	a        int
	function func(int)
}

func (c cmd) run() {
	c.function(c.a)
}

func SomeFunction(a int) {
	fmt.Println(a ^ 2)
}

func example_command() {
	var cmd1 ICmd = cmd{2, SomeFunction}
	cmd1.run()
}
