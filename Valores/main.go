package main

import "fmt"

func main(){

	var numero int = 10
	numero2 := 20
	fmt.Println(numero, numero2)
	
	var numeroEntero = 40
	var numeroDoble = 40.4
	
	resultado := numeroEntero + int(numeroDoble)
	fmt.Println(resultado)
	

	var nombre string = "Francisco"
	apellido := "Rocha"
	nombreCompleto := nombre + " " + apellido
	fmt.Println(nombreCompleto)
	
	// STRING Y NUMERO
	miNombre := "Francisco Rocha"
	miEdad := 25

	miInfo := fmt.Sprintf("Hola mi nombre es %s y mi edad es %d", miNombre, miEdad)

	fmt.Println(miInfo)

}

