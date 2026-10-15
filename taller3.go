package main

import (
	"fmt"
	"strings"
)

func main() {
	var texto string

	fmt.Println("Ingrese una frase:")
	fmt.Scanln(&texto)

	palabras := strings.Fields(texto)
	conteo := make(map[string]int)

	for _, palabra := range palabras {
		conteo[palabra]++
	}

	mayor := 0
	repetida := ""

	for palabra, cantidad := range conteo {
		fmt.Println(palabra, ":", cantidad)

		if cantidad > mayor {
			mayor = cantidad
			repetida = palabra
		}
	}

	fmt.Println("Palabra más repetida:", repetida, mayor)
}