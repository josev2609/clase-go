package main

import (
	"errors"
	"fmt"
)

type dispositivos interface {
	Encender() error
	Apagar() error
	Estatus() string
}

type Router struct {
	Nombre    string
	encendido bool
}

func (r *Router) Encender() error {
	if r.encendido {
		return errors.New("El router esta encendido")
	}

	r.encendido = true
	return nil

}

func (r *Router) Apagar() error {
	if r.encendido {
		return errors.New("El router esta apagado")
	}

	r.encendido = false
	return nil
}

func (r *Router) Estatus() string {
	if r.encendido {
		return r.Nombre + ":" + "encendido"
	} else {
		return r.Nombre + ":" + "apagado"
	}
}

type Servidor struct {
	Nombre string
	Activo bool
}

func (s *Servidor) Encender() error {
	if s.Activo {
		return errors.New("El serviodr esta activo")
	}

	s.Activo = true
	return nil
}

func (s *Servidor) Inactivo() error {
	if s.Activo {
		return errors.New("El servidor esta inactivo")
	}

	s.Activo = false
	return nil
}

func (s *Servidor) Estatus() string {
	if s.Activo {
		return s.Nombre + ":" + "activo"
	} else {
		return s.Nombre + ":" + "inactivo"
	}
}

func OperarDispositivo(d dispositivos) {
	fmt.Println("Estatus inicial", d.Estatus())

	err := d.Encender()

	if err != nil {
		fmt.Println("ERROR:", err)

	}

	fmt.Println("Despues de haber encendido:", d.Estatus())

	er := d.Apagar()

	if er != nil {
		fmt.Println("ERROR:", er)

	}

	fmt.Println("Despues de haber apagado:", d.Estatus())

}

func main() {

	router := &Router{
		Nombre: "Cisco",
	}

	servidor := &Servidor{
		Nombre: "Servidor web",
	}

	fmt.Println("*****Router*****")
	OperarDispositivo(router)

	fmt.Println("*****Servidor*****")
	OperarDispositivo(servidor)
}
