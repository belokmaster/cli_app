package main

import (
	"cli_app/cmd/app"
	"fmt"
	"os"
)

func main() {
	// Запускаем обработку команд
	if err := app.Execute(); err != nil {
		fmt.Println("Ошибка при выполнении команды:", err)
		os.Exit(1)
	}
}
