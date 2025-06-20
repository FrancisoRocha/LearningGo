/* Descripción:
Crea un programa que almacene 10 números enteros en un slice. Luego:
	1.	Usa un for para recorrer el slice.
	2.	Convierte los números en float64 y calcula el promedio.
	3.	Si el promedio es mayor o igual a 7.0, imprime "Buen promedio", si es menor, "Necesitas mejorar".

Extra: Usa un switch para imprimir si cada número es:
	•	Par
	•	Impar

Pistas:
	•	Usa if promedio >= 7.0
	•	Para saber si un número es par: n % 2 == 0 */

package main

import "fmt"

func main() {

	numerosEnteros := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// fmt.Println(numerosEnteros)

	suma := 0.0

	for i, enteros := range numerosEnteros {
		fmt.Printf("indice: %d, Valor %d\n", i, enteros)
		suma += float64(enteros)

		switch {
		case enteros%2 == 0:
			fmt.Println("Es Par")
		default:
			fmt.Println("Es Impar")
		}
	}

	promedio := suma / float64(len(numerosEnteros))

	if promedio >= 7.0 {
		fmt.Println("Buen promedio")
	} else {
		fmt.Println("Necesitas mejorar")
	}

	fmt.Printf("Promedio: %.2f\n", promedio)

}
