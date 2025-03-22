package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Отметить задачу как выполненную",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		tasks := loadTasks()
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("ID задачи должен быть числом")
			return
		}

		for i, task := range tasks {
			if task.ID == id {
				tasks[i].Done = true
				saveTasks(tasks)
				fmt.Printf("Задача %d выполнена!\n", id)
				return
			}
		}

		fmt.Println("Задача не найдена")
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
