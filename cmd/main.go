package main

import (
	"fmt"

	"code"
)

func main() {
	fmt.Println(code.GeneratePassword(8, 1, true, true, false))
	fmt.Println(code.GeneratePassword(12, 123, true, true, false))
	fmt.Println(code.GeneratePassword(12, 123, true, true, true))
	fmt.Println(code.GeneratePassword(8, 1, false, false, false))
}
