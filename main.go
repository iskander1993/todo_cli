package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/iskander1993/todo_cli/todo"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "list":
			todo.ListTasks()
		case "exit":
			fmt.Println("Выход...")
			return
		default:
			fmt.Println("Неизвестная команда. Доступные команды: list, exit")
		}
	}
}
