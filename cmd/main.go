package main

import (
	"fmt"

	"code"
)

func main() {
	var password string
	fmt.Println("** Hexlet password generator ***")
	fmt.Print("Для оценки надежности, введите ваш пароль:\n")
	fmt.Scan(&password)
	fmt.Println("--- Результаты проверки ---")
	fmt.Println(code.CheckPassword(password))
}
