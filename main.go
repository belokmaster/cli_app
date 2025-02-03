package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("base.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for {
		fmt.Println("Допустимые операции:")
		fmt.Println("1 - добавить новую задачу.")
		fmt.Println("q - выйти из программы.")
		var input string
		fmt.Scan(&input)

		switch input {
		case "q":
			fmt.Println("Завершение работы программы.")
			return
		case "1":
			fmt.Println("Выбрана команда добавления задачи в список.")
			fmt.Println("Пожалуйста, введите задачу.")

			var newTask string
			fmt.Scan(&newTask)

			_, err = file.WriteString(newTask + "\n")
			if err != nil {
				fmt.Println("Ошибка при записи в файл:", err)
				return
			}

			fmt.Println("Задача успешно добавлена в список.")
			return
		default:
			fmt.Println("Пожалуйста повторите ввод. Некорректная команда.")
		}
	}
}
