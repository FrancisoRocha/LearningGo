
🧩 Ejercicio 7: Clasificador de calificaciones

Objetivo: Practicar funciones, condicionales y slices en Go.

⸻

🔧 Instrucciones:
	1.	Crea una función llamada clasificarNota que reciba una calificación float64 y devuelva un string con la categoría:
	•	90 - 100: “Excelente”
	•	80 - 89: “Muy Bien”
	•	70 - 79: “Bien”
	•	60 - 69: “Suficiente”
	•	< 60: “Reprobado”
	2.	En la función main:
	•	Declara un slice de calificaciones, por ejemplo: []float64{95, 82, 74, 66, 48}
	•	Recorre el slice con un for y para cada calificación:
	•	Llama a la función clasificarNota
	•	Imprime: "La calificación X es: Y"

La calificación 95.00 es: Excelente  
La calificación 82.00 es: Muy Bien  
La calificación 74.00 es: Bien  
La calificación 66.00 es: Suficiente  
La calificación 48.00 es: Reprobado
