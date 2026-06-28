package code

import (
	"fmt"
	"strings"
	"unicode"
)

const lowercase = "abcdefghijklmnopqrstuvwxyz"
const uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const digits = "0123456789"
const special = "!@#$%^&*"

func NextRandom(number int) int {
	return (16807 * number) % 2147483647
}

func GeneratePassword(length, seed int, useUppercase, useDigits, useSpecial bool) string {
	// проверяем аргумент length
	if length <= 0 {
		return ""
	}
	// проверяем аргумент seed
	if seed < 0 {
		seed = -seed
	}
	if seed == 0 {
		seed = 1
	}
	current := seed
	result := ""
	alphabet := lowercase
	// создаём набор символов в зависимости от флагов
	if useUppercase {
		alphabet += uppercase
	}
	if useDigits {
		alphabet += digits
	}
	if useSpecial {
		alphabet += special
	}
	// проходим в цикле и генерируем пароль
	for range length {
		current = NextRandom(current)
		index := current % len(alphabet)
		result = result + string(alphabet[index])
	}
	// возвращаем результат
	return result
}

func CheckPassword(password string) string {
	score := 0
	verdict := ""
	// проверяем длину пароля
	if len(password) >= 8 {
		score += 1
	}
	// проверяем наличие строчной латинской буквы
	for _, s := range password {
		if strings.ContainsRune(lowercase, s) {
			score += 1
			break
		}
	}
	// проверяем наличие заглавной латинской буквы
	for _, s := range password {
		if strings.ContainsRune(uppercase, s) {
			score += 1
			break
		}
	}
	// проверяем наличие цифры
	for _, s := range password {
		if unicode.IsNumber(s) {
			score += 1
			break
		}
	}
	// проверяем наличие спецсимвола
	for _, s := range password {
		if strings.ContainsRune(special, s) {
			score += 1
			break
		}	
	}
	// проверяем количество баллов и выносим вердикт
	switch {
	case score <= 2:
		verdict = "Слабый"
	case score == 3:
		verdict = "Средний"
	case score == 4:
		verdict = "Надёжный"
	case score == 5:
		verdict = "Очень надёжный"
	}
	message := fmt.Sprintf("%s пароль (оценка %d из 5)", verdict, score)
	// возвращаем результат
	return message
}
