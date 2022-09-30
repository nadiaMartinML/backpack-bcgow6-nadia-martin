package main

import "fmt"

// ejercicio 1
var (
	miNombre  = "Nadia"
	direccion = "Urquiza"
)

// ejercicio 2
var temperatura float64 = 25.7
var humedad int = 80
var presion float64 = 60.7

// ejercicio 3
//var nombre string   //corregida, empezaba con un numero
//var apellido string // correcta
//var edad int        //corregida, estaba al reves, primero es el nombre de la variable y luego el tipo de dato
//apellido := 6  corregida, no puede empezar con un numero, funcionaría dentro de una funcion
//var licencia_de_conducir = true // correcta
//var estaturaDeLaPersona int     // corregida, el nombre de las variables no puede tener espacios
//cantidadDeHijos := 2 correcta, pero funcionaria dentro de una función

//ejercicio 4
//var apellido string = "Gomez" //correcta
//var edad int = 35 //corregida, los valores de tipo int no llevan comillas
//boolean := "false"; corregida, los valores de tipo boolean no llevan comillas
//var sueldo string = "45857.90" // corregida, los valores string llevan comillas, segunda solución abajo...
//var sueldo float64 = 45857.90 // corregida, al ser un valor númerico el tipo de valor debería de ser float64
//var nombre string = "Julián" // correcta

func main() {
	// ejercicio 1
	fmt.Println(miNombre, direccion)
	//ejercicio 2
	fmt.Println(temperatura, humedad, presion)
}
