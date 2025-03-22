package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Показать список задач",
	Run: func(cmd *cobra.Command, args []string) {
		tasks := loadTasks()

		if len(tasks) == 0 {
			fmt.Println("Нет задач.")
			return
		}

		for _, task := range tasks {
			status := " "
			if task.Done {
				status = "✓"
			}
			fmt.Printf("[%s] %d: %s\n", status, task.ID, task.Title)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
