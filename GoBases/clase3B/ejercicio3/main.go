package main

import "fmt"

// Una empresa nacional se encarga de realizar venta de productos, servicios y mantenimiento.
// Para ello requieren realizar un programa que se encargue de calcular el precio total de
// Productos, Servicios y Mantenimientos. Debido a la fuerte demanda y para optimizar la velocidad
// requieren que el cálculo de la sumatoria se realice en paralelo mediante 3 go routines.

// Se requieren 3 estructuras:
// Productos: nombre, precio, cantidad.
// Servicios: nombre, precio, minutos trabajados.
// Mantenimiento: nombre, precio.

// Se requieren 3 funciones:
// Sumar Productos: recibe un array de producto y devuelve el precio total (precio * cantidad).
// Sumar Servicios: recibe un array de servicio y devuelve el precio total (precio * media hora trabajada,
// 	si no llega a trabajar 30 minutos se le cobra como si hubiese trabajado media hora).
// Sumar Mantenimiento: recibe un array de mantenimiento y devuelve el precio total.

// Los 3 se deben ejecutar concurrentemente y al final se debe mostrar por pantalla el monto final (sumando el total de los 3).

type product struct {
	name     string
	price    float64
	quantity int
}

type service struct {
	name    string
	price   float64
	minutes int
}

type maintenance struct {
	name  string
	price float64
}

func addProduct(products ...product) float64 {
	totalPriceProduct := 0.0
	for _, product := range products {
		totalPriceProduct += (product.price * float64(product.quantity))
	}
	fmt.Println("El total de productos es: $", totalPriceProduct)
	return totalPriceProduct
}

func addService(services ...service) float64 {
	totalPriceService := 0.0
	for _, service := range services {
		if service.minutes > 30 {
			totalPriceService += (service.price * 30.0)
		}
		totalPriceService += (service.price * float64(service.minutes))
	}
	fmt.Println("El total de servicios es: $", totalPriceService)
	return totalPriceService
}

func addMaintenance(maintenances ...maintenance) float64 {
	totalPriceMaintenance := 0.0
	for _, maintenance := range maintenances {
		totalPriceMaintenance += maintenance.price
	}
	fmt.Println("El total de mantenimiento es: $", totalPriceMaintenance)
	return totalPriceMaintenance
}

func main() {
	products := []product{
		{name: "Producto 1", price: 100, quantity: 20},
		{name: "Producto 2", price: 200, quantity: 4},
		{name: "Producto 3", price: 300, quantity: 1},
		{name: "Producto 4", price: 200, quantity: 200},
	}

	services := []service{
		{name: "Programación", price: 200, minutes: 480},
		{name: "Paqueteria", price: 10234, minutes: 30},
		{name: "Encomienda", price: 300, minutes: 15},
		{name: "Limpieza", price: 100, minutes: 22},
	}

	maintenances := []maintenance{
		{name: "Mantenimiento 1", price: 400},
		{name: "Mantenimiento 2", price: 2500},
		{name: "Mantenimiento 3", price: 4100},
		{name: "Mantenimiento 4", price: 4100},
	}
}
