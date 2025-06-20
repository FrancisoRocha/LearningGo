// Ejemplo: alturaCM := 175 → alturaM := 1.75

package main

import "fmt"

func main(){

	alturaCM := 175
	alturaM := float64(alturaCM) / 100.0

	fmt.Printf("Mi altura es de %d cm lo que equivale a %.2f metros\n", alturaCM, alturaM)

}
