// Usa la constante math.Pi
// Importa el paquete "math"

package main

import (
	"fmt"
	"math"
)

func main(){

	var radio float64 = 10
	var area float64 = math.Pi * radio * radio

	fmt.Printf("El area del circulo con radio de %.2f es de %.2f\n", radio, area)
}
