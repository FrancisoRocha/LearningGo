//Usa for para sumar los números del 1 al 10. Imprime el resultado al final.


package main

import "fmt"


func main(){

	// Suma de los numero del 1 al 10
	suma := 0
	for i := 1; i <= 10; i++{
		suma += i
	}
	fmt.Println("La suma de los numero del 1 al 10 es: ", suma)

}

