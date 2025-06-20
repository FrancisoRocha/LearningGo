package main

import "fmt"

func main() {

	var arreglo [5]int
	fmt.Println("Arreglo completo", arreglo)

	arreglo[4] = 40
	fmt.Println("Arreglo completo", arreglo)

	fmt.Println("Arreglo en la pocision: ", arreglo[4])

	fmt.Println("El tamaño del Array es de: ", len(arreglo))

	lista := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println("Original:", lista)
	fmt.Println("El tamaño del Array es de: ", len(lista))

	lista2 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	fmt.Println("Inferido:", lista2)
	fmt.Println("El tamaño del Array es de: ", len(lista2))

}
