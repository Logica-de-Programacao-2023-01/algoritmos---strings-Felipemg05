package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Digite a primeira frase: ")
	scanner.Scan()
	frase1 := scanner.Text()

	fmt.Print("Digite a segunda frase: ")
	scanner.Scan()
	frase2 := scanner.Text()

	s1 := frase1 + frase2
	fmt.Println("A frase é ", s1)
}
