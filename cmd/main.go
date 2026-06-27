package main

import (
	"fmt"

	"code"
)

func main() {
	fmt.Println(code.GeneratePassword(30, true, true))
	fmt.Println(code.GeneratePassword(30, false, false))
}
