package main

import "fmt"

func main() {

	pesoKG := 70.0
	alturaCM := 175

	alturaMetros := float64(alturaCM) / 100.0

	imc := pesoKG / (alturaMetros * alturaMetros)

	fmt.Printf("Tu peso es de %.2f kg y tu altura es de %.2f m\n", pesoKG, alturaMetros)
	fmt.Printf("Tu peso IMC es de %.2f\n", imc)

	if imc < 18.5 {
		fmt.Println("Bajo Peso")
	} else if imc <= 24.9 {
		fmt.Println("Peso Normal")
	} else if imc <= 29.9 {
		fmt.Println("SobrePeso")
	} else if imc >= 30 {
		fmt.Println("Sufres de Obesidad")
	} else {
		fmt.Println("Error")
	}
}
