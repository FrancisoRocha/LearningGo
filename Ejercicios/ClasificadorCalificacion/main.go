
package main 

import "fmt"

func clasificarNota( nota float64) string {

	categoria := nota

	if categoria >= 90{
		return "Excelente"
	} else if categoria >= 80 {
		return "Muy Bien"
	} else if categoria >= 70{
		return "Bien"
	} else if categoria >= 60 {
		return "Suficiente"
	} else {
		return "Reprobado"
	} 
}

func main(){

	calificaciones := []float64{95, 82, 74, 66, 48}

	for _, calificacion := range calificaciones{
		fmt.Printf("La calificación %.2f es: %s\n", calificacion, clasificarNota(calificacion))
	}

}


