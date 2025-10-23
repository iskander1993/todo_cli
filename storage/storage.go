package storage

import (
	"encoding/json"
	"errors"
	"os"
)

// FilePath путь к JSON
const FilePath = "tasks.json"

// SaveTasks сохраняет список задач в файл
func SaveTasks[T any](tasks []T) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(FilePath, data, 0644)
}

// LoadTasks загружает список задач из файла
func LoadTasks[T any]() ([]T, error) {
	file, err := os.ReadFile(FilePath)
	if err != nil {

		if errors.Is(err, os.ErrNotExist) {
		}
		return []T{}, nil // если файла нет — возвращаем пустой срез
	}

	var tasks []T
	err = json.Unmarshal(file, &tasks)
	if err != nil {
		return []T{}, nil
	}
	return tasks, nil
}
