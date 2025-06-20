//Usa un bucle for para imprimir todos los números pares del 1 al 20.

package main

import "fmt"

func main() {

	i := 1

	for i <= 20 {
		if i%2 == 0 {
			fmt.Println("Numero: ", i)
		}
		i++
	}
}
