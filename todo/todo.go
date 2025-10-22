package todo

import (
	"fmt"
	"time"

	"github.com/iskander1993/todo_cli/storage"
)

type Task struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Done      bool   `json:"done"`
	CreatedAt string `json:"created_at"`
}

var tasks []Task
var maxID int

// Загружаем задачи из файла
func LoadTasksFromFile() {
	tasksFromFile, _ := storage.LoadTasks[Task]()
	tasks = tasksFromFile

	maxID = 0
	for _, task := range tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}
}

// Сохраняем задачи в файл
func SaveTasksToFile() {
	err := storage.SaveTasks(tasks)
	if err != nil {
		fmt.Println("Ошибка при сохранении задач:", err)
	}
}

// Добавление новой задачи
func AddTask(name string) {
	maxID++
	newTask := Task{
		ID:        maxID,
		Name:      name,
		Done:      false,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	tasks = append(tasks, newTask)
	SaveTasksToFile()
}

// Вывод списка задач
func ListTasks() {
	if len(tasks) == 0 {
		fmt.Println("Список задач пуст")
		return
	}
	for _, task := range tasks {
		status := "✖"
		if task.Done {
			status = "✔"
		}
		fmt.Printf("%d. [%s] %s\n", task.ID, status, task.Name)
	}
}

// Удаление задачи по ID
func RemoveTask(id int) {
	index := -1
	for i, task := range tasks {
		if task.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Println("Задача с таким ID не найдена")
		return
	}

	tasks = append(tasks[:index], tasks[index+1:]...)
	SaveTasksToFile()
}

// Обновление названия задачи
func UpdateTask(id int, newName string) {
	found := false
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Name = newName
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Задача с таким ID не найдена")
		return
	}

	SaveTasksToFile()
}

// Пометка задачи как выполненной
func MarkDone(id int) {
	found := false
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Done = true
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Задача с таким ID не найдена")
		return
	}

	SaveTasksToFile()
}
