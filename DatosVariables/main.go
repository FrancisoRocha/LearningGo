package main

import "fmt"

func main() {

	var nombrePersona string = "Francisco"
	var apellidoPersona = "Rocha"
	segundoNombre := "Angel"

	fmt.Println("Hola mundo, Soy " + nombrePersona)
	fmt.Println("Mi apellido es " + apellidoPersona)
	fmt.Println("Mi segundo nombre es " + segundoNombre)

	/* PARTE NUMERICA */
	var asoActual int16 = 2025
	// var asoIt int8 = 1234 --> Arroja error ya que no es numero valido
	var asoReducido int8 = 25
	edad := 25

	fmt.Println("El año actual es :", asoActual)
	fmt.Println("El año actual reducido es :", asoReducido)
	fmt.Println("Mi edad es :", edad)

	/* ARREGLOS */
	var listaFrutas = [5]string{"Mango", "Durazno"}
	// fmt.Println(listaFrutas)
	// fmt.Println(listaFrutas[0])
	fmt.Println(listaFrutas[1])

	// listaPaises := [4]string{"Mexico", "USA", "Rusia"}
	listaPaises := []string{"Mexico", "USA", "Rusia"}
	fmt.Println(listaPaises)
	// listaPaises[0] = "Colombia"
	listaPaises = append(listaPaises, "Chile")
	fmt.Println(listaPaises)

	listaPaises2 := listaPaises[1:3]
	fmt.Println(listaPaises2)

	listaPaises3 := listaPaises[2:]
	fmt.Println(listaPaises3)

	listaPaises4 := listaPaises[:4]
	fmt.Println(listaPaises4)

}
