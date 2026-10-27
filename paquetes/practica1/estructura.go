
package main
import "fmt"
type Producto struct {
ID int
Nombre string
Precio float64
Cantidad int
}
func NuevoProducto(id int,nombre string,precio float64,cantidad int) Producto {
return Producto{id,nombre,precio,cantidad}
}
func (p *Producto) ActualizarPrecio(precio float64) {
p.Precio=precio
}
func (p *Producto) ActualizarCantidad(cantidad int) {
p.Cantidad=cantidad
}
func (p Producto) MostrarInfo() {
fmt.Println("ID:",p.ID,"Nombre:",p.Nombre,"Precio:",p.Precio,"Cantidad:",p.Cantidad)
}
type Inventario struct {
Productos map[int]Producto
}
func NuevoInventario() Inventario {
return Inventario{make(map[int]Producto)}
}
func (i *Inventario) AgregarProducto(p Producto) {
i.Productos[p.ID]=p
}
func (i *Inventario) EliminarProducto(id int) {
delete(i.Productos,id)
}
func (i Inventario) BuscarProducto(id int) (Producto,bool) {
p,ok:=i.Productos[id]
return p,ok
}
func (i Inventario) ListarProductos() {
for _,p:=range i.Productos {
p.MostrarInfo()
}
}
func main() {
i:=NuevoInventario()
p:=NuevoProducto(1,"Laptop",500,10)
i.AgregarProducto(p)
p.ActualizarPrecio(600)
p.ActualizarCantidad(5)
i.AgregarProducto(p)
i.ListarProductos()
}
