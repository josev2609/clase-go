package vocales

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Ejecutar() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Ingrese una frase: ")
	frase, _ := reader.ReadString('\n')

	total := Contar(frase)

	fmt.Println("Cantidad de vocales:", total)
}

func Contar(frase string) int {

	var contador int

	frase = strings.ToLower(frase)

	for _, letra := range frase {

		switch letra {
		case 'a', 'á', 'e', 'é', 'i', 'í', 'o', 'ó', 'u', 'ú':
			contador++
		}
	}

	return contador
}
