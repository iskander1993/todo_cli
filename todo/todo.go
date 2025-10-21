package todo

import (
	"fmt"
	"github.com/iskander1993/todo_cli/storage"
)

func ListTasks() {
	tasks, err := storage.LoadTasks()
	if err != nil {
		fmt.Println("Ошибка при загрузке задач: ", err)
		return
	}
	if len(tasks) == 0 {
		fmt.Println("Список задач пуст")
		return
	}
	for i, task := range tasks {
		status := "✖"
		if done, ok := task["done"].(bool); ok && done {
			status = "✔"
		}
		name, _ := task["name"].(string)
		fmt.Printf("%d. [%s] %s\n", i+1, status, name)
	}
}
