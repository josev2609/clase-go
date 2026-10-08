package main

import "fmt"

func ContarVotos(){
	visitas:=make(map[string]int)
	visitas["inicio"]=10
	visitas["Noticias"]=20
	visitas["Deportes"]=25
	visitas["Galeria"]=20

	fmt.Println("Las noticias son:",visitas["Noticias"])
	for key, val:= range visitas{
		fmt.Println(key, ":", val)
	
	}

visitas["Galeria"]=20
fmt.Println("Las Noticias de Galeria son:",visitas["Galeria"])

visitas["Incio"]=200
fmt.Println("Las Noticias de Inicio son:",visitas["Inicio"])



}

func suma(map[string]int)int{
	visitas := make(map)
}