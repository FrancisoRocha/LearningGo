package main

import (
	"fmt"
	"strings"
)

func main() {

	separator := strings.Repeat("-", 20)

	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++
	}

	fmt.Println(separator)

	for numero := 0; numero < 5; numero++ {
		fmt.Println(numero)
	}

	fmt.Println(separator)

	for rango := range 5 {
		fmt.Println("rango: ", rango)
	}

	fmt.Println(separator)

	for {
		fmt.Println("Hola desde el loop")
		break
	}

	fmt.Println(separator)

	for valor := range 8 {
		if valor%2 == 0 {
			continue
		}
		fmt.Println(valor)
	}

}
