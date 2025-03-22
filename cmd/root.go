package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "Простой CLI-менеджер задач",
	Long:  "Менеджер задач на Go с использованием cobra.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Используйте 'todo --help' для списка команд")
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
