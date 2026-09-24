package main

import (
	"fmt"
	"practica/operaciones"
	"practica/saludo"
)

func main(){

	fmt.Println("Bienvenidos a la clase de paquetes")
	mensaje:=saludo.Saludar("Emilio")
	fmt.Println(mensaje)
	suma:=operaciones.Suma(5, 5 ) 
	fmt.Println(suma)
}
