package main

import (
	"fmt"
	"strings"
)

func main() {

	//Ejercicio 1
	//La Real Academia Española quiere saber cuántas letras tiene una palabra y
	//luego tener cada una de las letras por separado para deletrearla.
	//1.Crear una aplicación que tenga una variable con la palabra e imprimir la cantidad de letras que tiene la misma.
	//2.Luego imprimí cada una de las letras.

	palabra := "nan"
	cantidadLetras := len(palabra)
	fmt.Println(cantidadLetras)
	letras := strings.SplitAfter(palabra, "")
	fmt.Println(letras)

	for _, letra := range letras {
		fmt.Println(letra)
	}

	//Ejercicio 2
	//Un banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los mismos.
	//Para ello tiene ciertas reglas para saber a qué cliente se le puede otorgar.
	//Solo le otorga préstamos a clientes cuya edad sea mayor a 22 años, se encuentren empleados y
	//tengan más de un año de antigüedad en su trabajo. Dentro de los préstamos que otorga no les
	//cobrará interés a los que su sueldo es mejor a $100.000.
	//Es necesario realizar una aplicación que tenga estas variables y que imprima un mensaje de acuerdo a cada caso.
	//Tip: tu código tiene que poder imprimir al menos 3 mensajes diferentes.

	edad := 33
	esEmpleado := true
	antiguedad := 3
	sueldo := 140000

	if edad > 22 && esEmpleado && antiguedad > 1 {
		if sueldo > 100000 {
			fmt.Println("Se le otorgó el prestamo, no se le cobrará interés")
		} else {
			fmt.Println("Se le otorgó el préstamo.")
		}
	} else {
		fmt.Println("No se le otorgará el préstamo.")
	}

	//Ejercicio 3
	//Realizar una aplicación que contenga una variable con el número del mes.
	//1. Según el número, imprimir el mes que corresponda en texto.
	//2. ¿Se te ocurre si se puede resolver de más de una manera? ¿Cuál elegirías y por qué?
	//Ej: 7, Julio

	//SE PUEDE RESOLVER TAMBIEN CON UN IF ANIDADO

	var numMes int
	numMes = 7
	switch numMes {
	case 1:
		fmt.Println(numMes, "Enero")
	case 2:
		fmt.Println(numMes, "Febrero")
	case 3:
		fmt.Println(numMes, "Marzo")
	case 4:
		fmt.Println(numMes, "Abril")
	case 5:
		fmt.Println(numMes, "Mayo")
	case 6:
		fmt.Println(numMes, "Junio")
	case 7:
		fmt.Println(numMes, "Julio")
	case 8:
		fmt.Println(numMes, "Agosto")
	case 9:
		fmt.Println(numMes, "Septiembre")
	case 10:
		fmt.Println(numMes, "Octubre")
	case 11:
		fmt.Println(numMes, "Noviembre")
	case 12:
		fmt.Println(numMes, "Diciembre")
	}

	//Ejercicio 4
	//Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados.
	//Según el siguiente mapa, ayuda  a imprimir la edad de Benjamin.
	//var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	//Por otro lado también es necesario:
	// - Saber cuántos de sus empleados son mayores de 21 años.
	// - Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
	// - Eliminar a Pedro del mapa.

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	fmt.Println(employees["Benjamin"])
	old21 := 0
	for _, empl := range employees {
		if empl > 21 {
			old21++
		}
	}
	fmt.Println("Los mayores de 21 son en total: ", old21)
}
