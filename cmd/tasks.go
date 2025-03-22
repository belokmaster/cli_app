package cmd

import (
	"encoding/json"
	"os"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var taskFile = "tasks.json"

func loadTasks() []Task {
	var tasks []Task
	file, err := os.ReadFile(taskFile)
	if err == nil {
		json.Unmarshal(file, &tasks)
	}
	return tasks
}

func saveTasks(tasks []Task) {
	data, _ := json.MarshalIndent(tasks, "", "  ")
	os.WriteFile(taskFile, data, 0644)
}
