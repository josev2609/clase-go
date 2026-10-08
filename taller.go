package main

import "fmt"

func main() {

	notas := [6][4]float64{
		{8, 9, 7, 10},
		{6, 7, 8, 9},
		{10, 9, 10, 8},
		{7, 6, 8, 7},
		{9, 8, 9, 10},
		{5, 6, 7, 8},
	}

	var total float64

	for i := 0; i < 6; i++ {


		slice := notas[i][:]

		suma := 0.0
		mayor := slice[0]
		menor := slice[0]

		for _, nota := range slice {
			suma += nota

			if nota > mayor {
				mayor = nota
			}

			if nota < menor {
				menor = nota
			}
		}

		promedio := suma / 4
		total += suma

		fmt.Println("Estudiante", i+1)
		fmt.Println("Promedio:", promedio)
		fmt.Println("Nota mayor:", mayor)
		fmt.Println("Nota menor:", menor)
		fmt.Println()
	}

	fmt.Println("Promedio general:", total/24)
}