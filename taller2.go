package main

import "fmt"

func mayorVoto(votos map[string]int) string {

	mayor := 0
	actividad := ""

	for nombre, cantidad := range votos {
		if cantidad > mayor {
			mayor = cantidad
			actividad = nombre
		}
	}

	return actividad
}

func main() {

	votos := map[string]int{
		"deportes":    0,
		"videojuegos": 0,
		"cine":        0,
		"musica":      0,
	}

	for i := 0; i < 5; i++ {

		var voto string

		fmt.Print("Ingrese su voto (deportes, videojuegos, cine, musica): ")
		fmt.Scan(&voto)

		votos[voto]++
	}

	fmt.Println("Resultados:")

	for actividad, cantidad := range votos {
		fmt.Println(actividad, ":", cantidad, "votos")
	}

	ganador := mayorVoto(votos)

	fmt.Println("\nActividad con mayor aceptación:", ganador)
}