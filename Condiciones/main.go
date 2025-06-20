package main

import "fmt"

func main() {

	nombre := "Amin"
	edad := 25

	if nombre == "Amin" {
		fmt.Println("Hola Amin")
	} else {
		fmt.Println("Hola Desconocido")
	}

	// VALIDAR SI ES MAYOR DE EDAD
	if edad >= 18 {
		fmt.Println("Ya puedes votar")
	} else {
		fmt.Println("No puedes votar")
	}

	// VALIDAR SI ES NUMERO PAR O IMPAR
	if 8%2 == 0 {
		fmt.Println("El numero 8 es par")
	} else {
		fmt.Println("El numero 8 es impar")
	}

	// VALIDAR SI ES NEGATTIVO, DIGITO O MAYOR
	if numero := 9; numero < 0 {
		fmt.Println(numero, "Es negativo")
	} else if numero < 10 {
		fmt.Println(numero, "Es un digito")
	} else {
		fmt.Println(numero, "Es mayor a 9")
	}

	// NUMERO PAR O IMPAR
	numero1 := 11

	if numero1%2 == 0 {
		fmt.Println(numero1, "Numero PAR")

	} else {
		fmt.Println(numero1, "Numero es IMPAR")
	}

}
