package main

import (
	"fmt"
	"tallerMonedas/monedas"
	"tallerMonedas/vocales"
)

func main() {

	var opcion int

	fmt.Println("1. Conversor de monedas")
	fmt.Println("2. Contador de vocales")
	fmt.Print("Elija una opcion: ")
	fmt.Scan(&opcion)

	if opcion == 1 {
		monedas.Ejecutar()
	}

	if opcion == 2 {
		vocales.Ejecutar()
	}
}