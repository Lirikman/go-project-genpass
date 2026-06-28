package main

import (
	"fmt"
	"math/rand/v2"

	"code"
)

func main() {
	var answer string
	var password string
	fmt.Println("*** Hexlet password generator and strength check ***")
	fmt.Println("Я могу для Вас сгенерировать надеждный пароль!\nИли проверить Ваш пароль на надёжность.")
	fmt.Print("Для генерации введите - 'gen'\nДля проверки введите - 'check'\nВведите ответ: ")
	fmt.Scan(&answer)
	if answer != "gen" && answer != "check" {
		fmt.Println("Введён некорректный ответ!")
		return
	}
	if answer == "gen" {
		var length int
		fmt.Print("Введите желаемую длину пароля, но не менее 8: ")
		fmt.Scan(&length)
		if length < 8 {
			fmt.Println("Надёжный пароль содержит 'НЕ МЕНЕЕ' 8 символов!")
			return
		}
		seed := rand.IntN(10000)
		password := code.GeneratePassword(length, seed, true, true, true)
		fmt.Println("Ваш пароль: ", password)
	}
	if answer == "check" {
		fmt.Print("Для оценки надежности, введите ваш пароль:\n")
		fmt.Scan(&password)
		fmt.Println("--- Результаты проверки ---")
		fmt.Println(code.CheckPassword(password))
	}
}
