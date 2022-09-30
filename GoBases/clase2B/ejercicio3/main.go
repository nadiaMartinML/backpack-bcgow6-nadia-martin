package main

import "fmt"

//Varias tiendas de ecommerce necesitan realizar una funcionalidad en Go para administrar productos y
//retornar el valor del precio total.
//Las empresas tienen 3 tipos de productos:
//  Pequeño, Mediano y Grande. (Se espera que sean muchos más)
//Existen costos adicionales por mantener el producto en el almacén de la tienda, y costos de envío.
//Sus costos adicionales son:
//  Pequeño: El costo del producto (sin costo adicional)
//  Mediano: El costo del producto + un 3% por mantenerlo en existencia en el almacén de la tienda.
//  Grande: El costo del producto + un 6%  por mantenimiento, y un costo adicional  por envío de $2500.
//Requerimientos:
//  Crear una estructura 'tienda' que guarde una lista de productos.
//  Crear una estructura 'producto' que guarde el tipo de producto, nombre y precio
//  Crear una interface 'Producto' que tenga el método 'CalcularCosto'
//  Crear una interface 'Ecommerce' que tenga los métodos 'Total' y 'Agregar'.
//  Se requiere una función 'nuevoProducto' que reciba el tipo de producto, su nombre y precio y devuelva un Producto.
//  Se requiere una función 'nuevaTienda' que devuelva un Ecommerce.
// Interface Producto:
//  El método 'CalcularCosto' debe calcular el costo adicional según el tipo de producto.
// Interface Ecommerce:
// - El método 'Total' debe retornar el precio total en base al costo total de los productos y los adicionales si los hubiera.
// - El método 'Agregar' debe recibir un producto y añadirlo a la lista de la tienda

type product struct {
	productType string
	name        string
	price       float64
}

type store struct {
	productsList []product
}

type Producto interface {
	calculateCost(product) float64
}

type Ecommerce interface {
	total() float64
	add(product)
}

func (p *product) calculateCost() float64 {
	cost := 0.0
	switch p.productType {
	case "Pequeño":
		cost = p.price
	case "Mediano":
		cost = p.price + (p.price * 0.03)
	case "Grande":
		cost = p.price + (p.price * 0.06) + 2500
	}
	return cost
}

func (s *store) total() float64 {
	totalPrice := 0.0
	for _, product := range s.productsList {
		totalPrice += product.calculateCost()
	}
	return totalPrice
}

func (s *store) add(newProduct product) {
	s.productsList = append(s.productsList, newProduct)
}

func newProduct(productType, name string, price float64) product {
	return product{productType: productType, name: name, price: price}
}

func newStore() Ecommerce {
	return &store{}
}

func main() {
	p1 := newProduct("Pequeño", "Vaso", 600.0)
	p2 := newProduct("Mediano", "Cafetera", 2000.0)
	p3 := newProduct("Grande", "Televisor", 40000.0)
	s1 := store{}
	s1.add(p1)
	s1.add(p2)
	s1.add(p3)

	total := s1.total()
	fmt.Printf("El total del costo de la tienda es: $%v\n", total)
}
