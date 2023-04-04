package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var contador int
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Digite uma palavra : ")
	scanner.Scan()
	palavra := scanner.Text()

	fmt.Print("Digite um caractere : ")
	scanner.Scan()
	Caractere := scanner.Text()

	for i := 0; i < len(palavra); i++ {
		if string(palavra[i]) == Caractere {
			contador++
		}
	}
	fmt.Println("Contador=", contador)
}
