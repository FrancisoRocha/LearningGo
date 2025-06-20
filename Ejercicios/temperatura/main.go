// Fórmula: F = C * 9/5 + 32
// 1. Declara una variable float64 para grados Celsius
// 2. Convierte a Fahrenheit y muestra ambos valores

package main 

import "fmt"

func main(){

	var celsius float64 = 25
	var fahrenheit float64 = celsius * 9/5 + 32

	fmt.Printf("%.2f°C son %.2f°F\n", celsius, fahrenheit)

}

