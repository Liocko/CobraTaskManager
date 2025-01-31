package cmd

import (
	"encoding/json"
	"os"
)

type Task struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

var TaskList = []Task{
	{"TestTask", "Testing adding tasks", "✗"},
	{"SecondTask", "Another test task", "✓"},
}

const TaskFile = "tasks.json"

func saveTasks() error {
	data, err := json.Marshal(TaskList)
	if err != nil {
		return err
	}
	return os.WriteFile(TaskFile, data, 0644)
}

func loadTasks() error {
	if _, err := os.Stat(TaskFile); os.IsNotExist(err) {
		return nil
	}
	data, err := os.ReadFile(TaskFile)
	if err != nil {
		if os.IsNotExist(err) {
			TaskList = []Task{}
			return nil
		}
	}
	return json.Unmarshal(data, &TaskList)
}
