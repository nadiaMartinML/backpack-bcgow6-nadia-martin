package main

import (
	"fmt"
	"os"
	"strings"
)

//Ejercicio 1
//Una empresa que se encarga de vender productos de limpieza necesita:
// - Implementar una funcionalidad para guardar un archivo de texto, con la información de productos comprados,
//   separados por punto y coma (csv).
// - Debe tener el id del producto, precio y la cantidad.
// - Estos valores pueden ser hardcodeados o escritos en duro en una variable.

//Ejercicio 2
//La misma empresa necesita leer el archivo almacenado, para ello requiere que:
// se imprima por pantalla mostrando los valores tabulados, con un título
// (tabulado a la izquierda para el ID y a la derecha para el Precio y Cantidad),
// el precio, la cantidad y abajo del precio se debe visualizar el total (Sumando precio por cantidad)

type Product struct {
	id       int
	price    float64
	quantity int
}

func writeFile() {
	p1 := fmt.Sprintf("%d;%f;%d\n", 111223, 30012.00, 1)
	p2 := fmt.Sprintf("%d;%f;%d\n", 444321, 1000000.00, 4)
	p3 := fmt.Sprintf("%d;%f;%d\n", 434321, 50.50, 1)
	textInsert := fmt.Sprint(p1, p2, p3)

	err := os.WriteFile("products.csv", []byte(textInsert), 0644)
	if err != nil {
		panic("Hubo un error en la creación del archivo...")
	}
}

func main() {

	//writeFile()

	data, err := os.ReadFile("products.csv")
	if err != nil {
		panic("Hubo un error en la lectura del archivo...")
	}

	//var total float64
	fmt.Print("ID \t\t", "Precio \t\t\t", "Cantidad \n")
	dataString := strings.Split(string(data), "\n")
	for _, datos := range dataString {
		value := strings.Split(datos, ";")
		// price, err1 := strconv.ParseFloat(value[1], 64)
		// quantity, err2 := strconv.ParseFloat(value[2], 64)
		// if err1 == nil && err2 == nil {
		// 	total += price * quantity
		// }
		for _, v := range value {
			fmt.Printf("%v\t\t", v)
		}
		fmt.Printf("\n")
	}

	//fmt.Printf("\t\t\t%v", total)
}
