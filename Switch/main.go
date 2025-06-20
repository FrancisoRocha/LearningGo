package main

import (
	"fmt"
	"time"
)

func main() {

	i := 3
	fmt.Println("Escribe ", i, " como")

	switch i {
	case 1:
		fmt.Println("Uno")
	case 2:
		fmt.Println("Dos")
	case 3:
		fmt.Println("Tres")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Dias de Descanso")
	default:
		fmt.Println("A Trabajar :(")
	}

	tiempo := time.Now()
	fmt.Println("La hora es ", tiempo)

	switch {
	case tiempo.Hour() < 12:
		fmt.Println("Buenos Dias")
	case tiempo.Hour() < 18:
		fmt.Println("Buenas Tardes")
	default:
		fmt.Println("Buenas Noches")
	}

	// RETO CAMBIAR WEEKDARY A DIAS DE LA SEMANA EN ESPAÑOL
	switch time.Now().Weekday() {
	case time.Friday:
		fmt.Println("Viernes")
	case time.Saturday:
		fmt.Println("Sabado")
	case time.Sunday:
		fmt.Println("Domingo")
	default:
		fmt.Println("A Trabajar :(")
	}

}
