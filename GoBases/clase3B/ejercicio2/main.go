package main

import "fmt"

//Una importante empresa de ventas web necesita agregar una funcionalidad para agregar productos a los usuarios.
//Para ello requieren que tanto los usuarios como los productos tengan la misma dirección de memoria en el main
//del programa como en las funciones.
//Se necesitan las estructuras:
//Usuario: Nombre, Apellido, Correo, Productos (array de productos).
//Producto: Nombre, precio, cantidad.
//Se requieren las funciones:
//Nuevo producto: recibe nombre y precio, y retorna un producto.
//Agregar producto: recibe usuario, producto y cantidad, no retorna nada, agrega el producto al usuario.
//Borrar productos: recibe un usuario, borra los productos del usuario.

type user struct {
	name         string
	lastName     string
	mail         string
	productsList []product
}

type product struct {
	name     string
	price    float64
	quantity int
}

func newProduct(name *string, price *float64) *product {
	return &product{name: *name, price: *price}
}

func (u *user) addProduct(product *product, quantity *int) {
	product.quantity = *quantity
	u.productsList = append(u.productsList, *product)
}

func (u *user) deleteProduct() {
	u.productsList = []product{}
}

func main() {
	u1 := &user{
		name:     "nani",
		lastName: "martin",
		mail:     "nmm@mail.com",
	}
	fmt.Printf("El usuario %v %v con mail %v tiene los siguientes productos:\n%v\n", u1.name, u1.lastName, u1.mail, u1.productsList)

	var (
		name1     string  = "vaso"
		price1    float64 = 80.0
		quantity1 int     = 4
		name2     string  = "lapiz"
		price2    float64 = 13
		quantity2 int     = 2
	)
	p1 := newProduct(&name1, &price1)
	p2 := newProduct(&name2, &price2)

	u1.addProduct(p1, &quantity1)
	u1.addProduct(p2, &quantity2)
	fmt.Printf("El usuario %v %v con mail %v tiene los siguientes productos:\n%v\n", u1.name, u1.lastName, u1.mail, u1.productsList)

	u1.deleteProduct()
	fmt.Printf("El usuario %v %v con mail %v tiene los siguientes productos:\n%v\n", u1.name, u1.lastName, u1.mail, u1.productsList)

}
