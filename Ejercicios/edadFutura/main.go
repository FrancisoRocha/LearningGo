// Ejemplo: edadActual := 22, añosDespués := 5 → "Tendrás 27 años"

package main

import "fmt"

func main(){

	var edadActual int = 25
	var edadFutura int = 5 + edadActual

	fmt.Printf("Mi edad es %d y en 5 años tendre %d años", edadActual, edadFutura)

}