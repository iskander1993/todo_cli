package storage

import (
	"encoding/json"
	"os"
)

const FilePath = "tasks.json"

func LoadTasks() ([]map[string]interface{}, error) {
	file, err := os.ReadFile(FilePath)
	if err != nil {
		return []map[string]interface{}{}, nil
	}
	var tasks []map[string]interface{}
	err = json.Unmarshal(file, &tasks)
	if err != nil {
		return []map[string]interface{}{}, nil
	}
	return tasks, nil
}

func SaveTasks(tasks []map[string]interface{}) error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(FilePath, data, 0644)
}
