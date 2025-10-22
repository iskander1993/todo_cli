package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/iskander1993/todo_cli/todo"
)

func main() {
	// Загружаем задачи при старте программы
	todo.LoadTasksFromFile()

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Привет! Это To-Do CLI. Введите 'help' для списка команд.")

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch {
		// Список задач
		case input == "list" || input == "l":
			todo.ListTasks()

		// Добавление задачи
		case strings.HasPrefix(input, "add ") || strings.HasPrefix(input, "a "):
			name := strings.TrimPrefix(input, "add ")
			name = strings.TrimPrefix(name, "a ")
			if name == "" {
				fmt.Println("Введите название задачи после команды add")
				continue
			}
			todo.AddTask(name)
			fmt.Println("Задача добавлена!")

		// Удаление задачи
		case strings.HasPrefix(input, "remove ") || strings.HasPrefix(input, "r "):
			idStr := strings.TrimPrefix(input, "remove ")
			idStr = strings.TrimPrefix(idStr, "r ")
			id, err := strconv.Atoi(strings.TrimSpace(idStr))
			if err != nil {
				fmt.Println("Неверный ID")
				continue
			}
			todo.RemoveTask(id)
			fmt.Println("Задача удалена!")

		// Обновление названия задачи
		case strings.HasPrefix(input, "update ") || strings.HasPrefix(input, "u "):
			args := strings.Fields(strings.TrimPrefix(input, "update "))
			if len(args) < 2 {
				args = strings.Fields(strings.TrimPrefix(input, "u "))
			}
			if len(args) < 2 {
				fmt.Println("Использование: update <ID> <Новое название>")
				continue
			}
			id, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println("Неверный ID")
				continue
			}
			newName := strings.Join(args[1:], " ")
			todo.UpdateTask(id, newName)
			fmt.Println("Задача обновлена!")

		// Пометить задачу как выполненную
		case strings.HasPrefix(input, "done ") || strings.HasPrefix(input, "d "):
			idStr := strings.TrimPrefix(input, "done ")
			idStr = strings.TrimPrefix(idStr, "d ")
			id, err := strconv.Atoi(strings.TrimSpace(idStr))
			if err != nil {
				fmt.Println("Неверный ID")
				continue
			}
			todo.MarkDone(id)
			fmt.Println("Задача выполнена!")

		// Выход
		case input == "exit":
			fmt.Println("Выход...")
			return

		// Помощь
		case input == "help" || input == "h":
			fmt.Println(`Доступные команды:
list или l                     - показать список задач
add "название" или a           - добавить задачу
remove <ID> или r              - удалить задачу
update <ID> "новое название" или u - изменить название задачи
done <ID> или d                - пометить задачу выполненной
help или h                     - показать список команд
exit                            - выйти из программы`)

		// Неизвестная команда
		default:
			fmt.Println("Неизвестная команда. Введите 'help' для списка команд.")
		}
	}
}
