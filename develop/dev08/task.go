package main

import (
	"L2/develop/dev08/internal"
	"L2/develop/dev08/internal/builtins"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	ReadCmd()
}

func ReadCmd() {
	var (
		commands []internal.ICommand
		paths    []string
	)
	env := os.Getenv("PATH")
	if env != "" {
		paths = strings.Split(env, ":")
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		if line, _, err := reader.ReadLine(); err != nil {
			log.Fatal(err)
		} else if string(line) == "\\quit" {
			break
		} else if commands, err = builtins.CreateCommands(string(line), paths); err != nil {
			log.Fatal(err)
		}
		internal.Execute(commands)
	}
}
