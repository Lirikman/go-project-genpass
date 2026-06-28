package main

import (
	"fmt"

	"code"
)

func main() {
	fmt.Println(code.CheckPassword("abc"))
	fmt.Println(code.CheckPassword("abcdefgh"))
	fmt.Println(code.CheckPassword("abcdef1234"))
	fmt.Println(code.CheckPassword("Abcdef1234"))
	fmt.Println(code.CheckPassword("Abcdef123!"))
}
