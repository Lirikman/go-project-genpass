package code

const lowercase = "abcdefghijklmnopqrstuvwxyz"
const uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const digits = "0123456789"

func GeneratePassword(length int, useUppercase, useDigits bool) string {
	alphabet := lowercase
	if useUppercase && !useDigits {
		alphabet += uppercase
	}
	if !useUppercase && useDigits {
		alphabet += digits
	}
	if useUppercase && useDigits {
		alphabet += uppercase
		alphabet += digits
	}
	result := ""
	for i := range length {
		index := i % len(alphabet)
		result = result + string(alphabet[index])
	}
	return result
}
