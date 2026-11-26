package main

import "fmt"

type InformacionProducto interface {
	MostrarInfo() string
}

type ProductoBase struct {
	ID       int
	Nombre   string
	Precio   float64
	Cantidad int
}

func (p ProductoBase) MostrarInfo() string {
	return fmt.Sprintf("ID: %d\nNombre: %s\nPrecio: %.2f\nCantidad: %d",
		p.ID, p.Nombre, p.Precio, p.Cantidad)
}

type Electronico struct {
	ProductoBase
	GarantiaMeses int
}

func NuevoProductoElectronico(id int, nombre string, precio float64, cantidad, garantia int) Electronico {
	return Electronico{ProductoBase{id, nombre, precio, cantidad}, garantia}
}

func (e Electronico) MostrarInfo() string {
	return e.ProductoBase.MostrarInfo() + fmt.Sprintf("\nGarantía: %d meses", e.GarantiaMeses)
}

type Alimento struct {
	ProductoBase
	FechaExpiracion string
}

func NuevoProductoAlimento(id int, nombre string, precio float64, cantidad int, fecha string) Alimento {
	return Alimento{ProductoBase{id, nombre, precio, cantidad}, fecha}
}

func (a Alimento) MostrarInfo() string {
	return a.ProductoBase.MostrarInfo() + "\nExpiración: " + a.FechaExpiracion
}

type Inventario struct {
	Productos []InformacionProducto
}

func (i *Inventario) AgregarProducto(p InformacionProducto) {
	i.Productos = append(i.Productos, p)
}

func (i Inventario) ListarProductos() {
	for _, p := range i.Productos {
		fmt.Println("\n" + p.MostrarInfo())

	}
}

func main() {
	inventario := Inventario{}
	var opcion int

	for {
		fmt.Println("\n1. Agregar electrónico")
		fmt.Println("2. Agregar alimento")
		fmt.Println("3. Listar productos")
		fmt.Println("4. Salir")
		fmt.Print("Opción: ")
		fmt.Scan(&opcion)

		if opcion == 1 {
			var id, cantidad, garantia int
			var nombre string
			var precio float64

			fmt.Print("ID, nombre, precio, cantidad y garantía: ")
			fmt.Scan(&id, &nombre, &precio, &cantidad, &garantia)

			inventario.AgregarProducto(
				NuevoProductoElectronico(id, nombre, precio, cantidad, garantia))

		} else if opcion == 2 {
			var id, cantidad int
			var nombre, fecha string
			var precio float64

			fmt.Print("ID, nombre, precio, cantidad y fecha: ")
			fmt.Scan(&id, &nombre, &precio, &cantidad, &fecha)

			inventario.AgregarProducto(
				NuevoProductoAlimento(id, nombre, precio, cantidad, fecha))

		} else if opcion == 3 {
			inventario.ListarProductos()

		} else if opcion == 4 {
			break
		} else {
			fmt.Println("Opción incorrecta")
		}
	}
}
