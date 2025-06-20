package main

import (
	"fmt"
	"slices"
)

func main() {

	var arregloCadenas []string
	fmt.Println("dato", "contenido", arregloCadenas, "condicion:", arregloCadenas == nil, "longitud:", len(arregloCadenas) == 0)

	arregloCadenas = make([]string, 3)
	arregloCadenas[0] = "A"
	arregloCadenas[1] = "B"
	arregloCadenas[2] = "C"

	fmt.Println("Datos:", arregloCadenas)
	fmt.Println("Dato especifico:", arregloCadenas[1])
	fmt.Println("Longitud:", len(arregloCadenas))

	// Agregar un elemento al final del arreglo
	arregloCadenas = append(arregloCadenas, "D")
	arregloCadenas = append(arregloCadenas, "E", "F")
	fmt.Println(arregloCadenas)
	fmt.Println("Longitud:", len(arregloCadenas))

	segundoArreglo := []string{"G", "H", "I"}
	fmt.Println(segundoArreglo)

	tercerArreglo := []string{"G", "H", "I"}
	fmt.Println(tercerArreglo)

	if slices.Equal(segundoArreglo, tercerArreglo) {
		fmt.Println("2do Arreglo es igual a 3er Arreglo")
	}
}
