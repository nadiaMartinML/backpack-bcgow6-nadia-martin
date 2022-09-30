package main

import (
	"errors"
	"fmt"
)

//Ejercicio 1
//Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar el sueldo,
//para cumplir el objetivo es necesario crear una función que devuelva el impuesto de un salario.
//Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17% del sueldo y
//si gana más de $150.000 se le descontará además un 10%.

func calcularImpuestos(sueldo float64) float64 {
	var impuesto float64
	if sueldo >= 50000 {
		impuesto = sueldo * 0.17
		return impuesto
	} else if sueldo >= 150000 {
		impuesto = sueldo * 0.10
		return impuesto
	}
	return 0
}

//Ejercicio 2
//Un colegio necesita calcular el promedio (por alumno) de sus calificaciones.
//Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y
//devuelva el promedio y un error en caso que uno de los números ingresados sea negativo

func calcularPromedio(calificaciones ...int) (promedio float64, err error) {
	sumaCalificaciones := 0
	for _, value := range calificaciones {
		if value < 0 {
			return 0, errors.New("No puede haber calificaciones en negativo.")
		}
		sumaCalificaciones += value
	}
	promedio = float64(sumaCalificaciones) / float64(len(calificaciones))
	return promedio, nil
}

//Ejercicio 3
//Una empresa marinera necesita calcular el salario de sus empleados basándose en
//la cantidad de horas trabajadas por mes y la categoría.
// - Si es categoría C, su salario es de $1.000 por hora
// - Si es categoría B, su salario es de $1.500 por hora más un %20 de su salario mensual
// - Si es de categoría A, su salario es de $3.000 por hora más un %50 de su salario mensual
//Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados
//por mes y la categoría, y que devuelva su salario.

func calculateSalary(minutes float64, category string) float64 {
	var salary float64
	if category == "C" {
		salary = (minutes * 1000) / 60
	}
	if category == "B" {
		salary = ((minutes * 1500) / 60) * 0.8
	}
	if category == "A" {
		salary = ((minutes * 3000) / 60) * 0.5
	}
	return salary
}

//Ejercicio 4
//Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas de calificaciones
//de los alumnos de un curso, requiriendo calcular los valores mínimo, máximo y promedio de sus calificaciones.
//Se solicita generar una función que indique qué tipo de cálculo se quiere realizar (mínimo, máximo o promedio)
//y que devuelva otra función ( y un mensaje en caso que el cálculo no esté definido) que se le puede pasar
//una cantidad N de enteros y devuelva el cálculo que se indicó en la función anterior

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func operation(operator string) (myFunc func(grades ...int) float64, err error) {
	switch operator {
	case "minimum":
		return minFunc, nil
	case "average":
		return averageFunc, nil
	case "maximum":
		return maxFunc, nil
	}
	return nil, errors.New("El calculo no esta definido")
}

func minFunc(grades ...int) float64 {
	minGrade := grades[0]
	for i := 1; i < len(grades); i++ {
		if grades[i] < minGrade {
			minGrade = grades[i]
		}
	}
	return float64(minGrade)
}

func averageFunc(grades ...int) float64 {
	sumGrades := 0
	for i := 0; i < len(grades); i++ {
		sumGrades += grades[i]
	}
	result := sumGrades / len(grades)
	return float64(result)
}

func maxFunc(grades ...int) float64 {
	maxGrade := grades[0]
	for i := 1; i < len(grades); i++ {
		if grades[i] > maxGrade {
			maxGrade = grades[i]
		}
	}
	return float64(maxGrade)
}

//Ejercicio 5
//Un refugio de animales necesita calcular cuánto alimento debe comprar para las mascotas.
//Por el momento solo tienen tarántulas, hamsters, perros, y gatos, pero se espera que puedan haber
//muchos más animales que refugiar.
// 1- perro necesitan 10 kg de alimento
// 2- gato 5 kg
// 3- Hamster 250 gramos.
// 4- Tarántula 150 gramos.
//Se solicita:
//Implementar una función Animal que reciba como parámetro un valor de tipo texto con el animal
//especificado y que retorne una función y un mensaje (en caso que no exista el animal)
//Una función para cada animal que calcule la cantidad de alimento en base a la cantidad del tipo de animal especificado.

const (
	dog       = "dog"
	cat       = "cat"
	hamster   = "hamster"
	tarantula = "tarantula"
)

func Animal(typeAnimal string) (func(int) float64, string) {
	switch typeAnimal {
	case "dog":
		return animalDog, ""
	case "cat":
		return animalCat, ""
	case "hamster":
		return animalHamster, ""
	case "tarantula":
		return animalTarantula, ""
	default:
		return nil, "El animal no está cargado...."
	}
}

func animalDog(cant int) float64 {
	return float64(10 * cant)
}

func animalCat(cant int) float64 {
	return float64(5 * cant)
}

func animalHamster(cant int) float64 {
	return 0.25 * float64(cant)
}

func animalTarantula(cant int) float64 {
	return 0.15 * float64(cant)
}

func main() {
	//Ejercicio 1
	fmt.Println(calcularImpuestos(90000))
	fmt.Println(calcularImpuestos(40000))
	fmt.Println(calcularImpuestos(160000))

	//Ejercicio 2
	fmt.Println(calcularPromedio(2, 5, 8, 10, 4))
	fmt.Println(calcularPromedio(2, 5, -8, 10, 4))

	//Ejercicio 3
	fmt.Println(calculateSalary(120, "C"))
	fmt.Println(calculateSalary(120, "B"))
	fmt.Println(calculateSalary(120, "A"))

	//Ejercicio 4
	minFunc, err := operation("minimum")
	if err != nil {
		return
	}
	averageFunc, err := operation("average")
	maxFunc, err := operation("maximum")

	minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	fmt.Println("El valor minimo es:", minValue)
	fmt.Println("El valor promedio es:", averageValue)
	fmt.Println("El valor maximo es:", maxValue)

	//Ejercicio 5
	animalDog, msg := Animal(dog)
	if msg != "" {
		fmt.Println(msg)
	}
	animalCat, msg := Animal(cat)
	if msg != "" {
		fmt.Println(msg)
	}
	animalHamster, msg := Animal(hamster)
	if msg != "" {
		fmt.Println(msg)
	}
	animalTarantula, msg := Animal(tarantula)
	if msg != "" {
		fmt.Println(msg)
	}

	var amount float64
	amount += animalDog(5)
	amount += animalCat(2)
	amount += animalHamster(4)
	amount += animalTarantula(8)

	fmt.Println("El valor a comprar es:", amount)

}
