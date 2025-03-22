package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Добавить новую задачу",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		title := args[0]
		tasks := loadTasks()

		newTask := Task{
			ID:    len(tasks) + 1,
			Title: title,
			Done:  false,
		}

		tasks = append(tasks, newTask)
		saveTasks(tasks)

		fmt.Printf("Добавлена задача: %s\n", title)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
