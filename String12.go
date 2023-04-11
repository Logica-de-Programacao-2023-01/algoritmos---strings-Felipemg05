package main

import (
	"fmt"
	"strings"
)

func main() {
	var x string
	fmt.Print("Escreva uma string:")
	fmt.Scan(&x)

	letras := strings.Split(x, "")

	if len(letras)%2 != 0 {
		palindromo := true
		for i := 0; i < len(letras)/2; i++ {
			if letras[i] != letras[len(letras)-i] {
				palindromo = false
				break
			}
		}
		if palindromo {
			fmt.Printf()
		}
	} else {
		fmt.Printf("A string %x não é um palindromo", x)
	}
}
