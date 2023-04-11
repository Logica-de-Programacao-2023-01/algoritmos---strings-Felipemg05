package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var x string
	fmt.Print("Escreva uma string:")
	fmt.Scan(&x)

	_, err := strconv.Atoi(x)
	if err != nil {
		fmt.Printf("A string %s não é uma sequencia numerica")
	} else {
		numeros := strings.Split(x, "")
		decrescente := true
		for i := 1; i < len(numeros); i++ {
			anterior
		}
	}
}
