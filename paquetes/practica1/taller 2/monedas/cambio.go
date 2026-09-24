package monedas

import "fmt"

func Ejecutar() {
	var dolares float64
	var moneda string

	fmt.Print("ingrese dolares ")
	fmt.Scan(&dolares)

	fmt.Print("Ingrese moneda EUR, LB, WON o BTC")
	fmt.Scan(&moneda)

	resultado := Convertir(dolares, moneda)

	fmt.Println("resultado", resultado)
}

func Convertir(dolares float64, moneda string) float64 {
	switch moneda {
	case "EUR":
		return dolares * 0.85
	case "LB":
		return dolares * 0.74
	case "WON":
		return dolares * 1400
	case "BTC":
		return dolares * 0.000009
	default:
		return 0
	}
}
