package code

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
