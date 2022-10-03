package main

import "fmt"

/*
Ejercicio 2 - Ecommerce
Una importante empresa de ventas web necesita agregar una funcionalidad para agregar productos a los
usuarios. Para ello requieren que tanto los usuarios como los productos tengan la misma dirección de
memoria en el main del programa como en las funciones.
Se necesitan las estructuras:
Usuario: Nombre, Apellido, Correo, Productos (array de productos).
Producto: Nombre, precio, cantidad.
Se requieren las funciones:
Nuevo producto: recibe nombre y precio, y retorna un producto.
Agregar producto: recibe usuario, producto y cantidad, no retorna nada, agrega el producto al usuario.
Borrar productos: recibe un usuario, borra los productos del usuario.
*/

type Usuario struct {
	NombreCompleto string
	Email          string
	Productos      []Producto
}

type Producto struct {
	Nombre   string
	precio   float64
	cantidad int
}

func nuevoProducto(nombre *string, precio *float64) *Producto {
	return &Producto{Nombre: *nombre, precio: *precio}
}

func (u *Usuario) agregarProducto(producto *Producto, cantidad *int) {
	u.Productos = append(u.Productos, *producto)
	producto.cantidad = *cantidad
}

func (u *Usuario) eliminarProducto() {
	u.Productos = nil
}

func main() {

	newUser := &Usuario{
		NombreCompleto: "Juan Perez",
		Email:          "juanp@email.com",
	}

	var nombre string = "arroz"
	var precio float64 = 56.0
	var cantidad int = 5

	newProduct := nuevoProducto(&nombre, &precio)
	newUser.agregarProducto(newProduct, &cantidad)

	fmt.Println(newProduct)
	fmt.Println(newUser.Productos)
}
