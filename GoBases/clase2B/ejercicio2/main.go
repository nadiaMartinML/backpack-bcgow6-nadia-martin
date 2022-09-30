package main

import "fmt"

//Ejercicio 2
//Una empresa de inteligencia artificial necesita tener una funcionalidad para crear una estructura que
//represente una matriz de datos.
//Para ello requieren una estructura Matrix que tenga los métodos:
//  Set:  Recibe una serie de valores de punto flotante e inicializa los valores en la estructura Matrix
//  Print: Imprime por pantalla la matriz de una formas más visible (Con los saltos de línea entre filas)
//La estructura Matrix debe contener los valores de la matriz, la dimensión del alto, la dimensión del ancho,
//si es cuadrática y cuál es el valor máximo.

type Matrix struct {
	values   []float64
	height   int
	width    int
	maxValue float64
}

func (m *Matrix) setMatrix(values ...float64) {
	m.values = values
}

func (m *Matrix) printMatrix() {
	for i := 0; i < m.height; i++ {
		for value := 0; value < m.width; value++ {
			fmt.Printf("%v\t", m.values[i*m.width+value])
		}
		fmt.Printf("\n")
	}
}

func (m *Matrix) maxValueFunc(values ...float64) float64 {
	m.maxValue = m.values[0]
	for i := 1; i < len(m.values); i++ {
		if m.values[i] > m.maxValue {
			m.maxValue = m.values[i]
		}
	}
	return m.maxValue
}

func main() {
	matrix := Matrix{height: 2, width: 3}
	matrix.setMatrix(4.9, 3.2, 2, 8.9, 10, 2.6)
	matrix.printMatrix()
	fmt.Println("El valor máximo de la matrix es:", matrix.maxValueFunc())
}
