package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Удалить задачу",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		tasks := loadTasks()
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("ID задачи должен быть числом")
			return
		}

		newTasks := []Task{}
		found := false
		for _, task := range tasks {
			if task.ID != id {
				newTasks = append(newTasks, task)
			} else {
				found = true
			}
		}

		if found {
			saveTasks(newTasks)
			fmt.Printf("Задача %d удалена\n", id)
		} else {
			fmt.Println("Задача не найдена")
		}
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
