package main

import "fmt"

type Producto struct {
	nombre   string
	precio   float64
	cantidad int
}

func mostrar(p Producto) {
	fmt.Println("Nombre:", p.nombre)
	fmt.Println("Precio:", p.precio)
	fmt.Println("Cantidad:", p.cantidad)
	fmt.Println("Total:", p.precio*float64(p.cantidad))

	if p.cantidad > 0 {
		fmt.Println("Disponible")
	} else {
		fmt.Println("Sin stock")
	}
}

func main() {
	p1 := Producto{"laptop", 800, 5}
	p2 := Producto{"mouse", 20, 10}
	p3 := Producto{"teclado", 30, 0}

	mostrar(p1)
	mostrar(p2)
	mostrar(p3)
}