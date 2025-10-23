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
	tasksFromFile, err := storage.LoadTasks[Task]()
	if err != nil {
		fmt.Println("⚠️  Ошибка при загрузке задач:", err)
	}
	tasks = tasksFromFile

	maxID = 0
	for _, task := range tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}
	if len(tasks) > 0 {
		fmt.Printf("✅ Загружено задач: %d\n", len(tasks))
	}
}

// Сохраняем задачи в файл
func SaveTasksToFile() {
	err := storage.SaveTasks(tasks)
	if err != nil {
		fmt.Println("❌ Ошибка при сохранении задач:", err)
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

// Вывод списка задач с фильтром
// filter: "all" - все, "done" - выполненные, "pending"
func ListTasks(filter string) {
	if len(tasks) == 0 {
		fmt.Println("📋 Список задач пуст")
		return
	}

	//Фильтруем задачи
	filteredTasks := []Task{}
	for _, task := range tasks {
		if filter == "all" ||
			(filter == "done" && task.Done) ||
			(filter == "pending" && !task.Done) {
			filteredTasks = append(filteredTasks, task)
		}
	}

	if len(filteredTasks) == 0 {
		switch filter {
		case "done":
			fmt.Println("📋 Нет выполненных задач")
		case "pending":
			fmt.Println("📋 Нет выполненных задач")
		}
		return
	}

	//Красивывй заголовок с статикой
	switch filter {
	case "done":
		fmt.Printf("\n✅ Выполненные задачи (%d):\n", len(filteredTasks))
	case "pending":
		fmt.Printf("\n⏳ Невыполненные задачи (%d):\n", len(filteredTasks))
	default:
		fmt.Printf("\n📋 Все задачи (всего: %d, выполнено: %d, осталось: %d):\n",
			len(tasks), countDone(), len(tasks)-countDone())
	}
	//Вывод задач

	for _, task := range filteredTasks {
		status := "✖"
		if task.Done {
			status = "✔"
		}
		fmt.Println("  %d. [%s] %s\n", task.ID, status, task.Name)

	}
	fmt.Println()
}

func countDone() int {
	count := 0
	for _, task := range tasks {
		if task.Done {
			count++
		}
	}
	return count
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
		fmt.Println("❌ Задача с таким ID не найдена")
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
		fmt.Println("❌ Задача с таким ID не найдена")
		return
	}

	SaveTasksToFile()
}

// Пометка задачи как выполненной
func MarkDone(id int) {
	found := false
	for i, task := range tasks {
		if task.ID == id {
			if !tasks[i].Done {
				fmt.Println("ℹ️Задача уже выполнена")
				return
			}
			tasks[i].Done = true
			found = true
			break
		}
	}

	if !found {
		fmt.Println("❌Задача с таким ID не найдена")
		return
	}

	SaveTasksToFile()
}
