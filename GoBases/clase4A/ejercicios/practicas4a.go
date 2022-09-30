package main

import (
	"errors"
	"fmt"
)

//Ejercicio 1
//En tu función 'main', define una variable llamada 'salary' y asignarle un valor de tipo 'int'.
//Crea un error personalizado con un struct que implemente 'Error()' con el mensaje 'error: el salario
//ingresado no alcanza el mínimo imponible' y lánzalo en caso de que 'salary' sea menor a 150.000.
//Caso contrario, imprime por consola el mensaje 'Debe pagar impuesto'.

var salary int

type myCustomError struct {
	message string
}

func (e *myCustomError) Error() string {
	return e.message
}

func payTax1(salary int) (string, error) {
	if salary < 150000 {
		return "", &myCustomError{
			message: "error: el salario ingresado no alcanza el mínimo imponible",
		}
	}
	return "Debe pagar impuesto.", nil
}

//Ejercicio 2
//Haz lo mismo que en el ejercicio anterior pero reformulando el código para que,
//en reemplazo de 'Error()', se implemente 'errors.New()'.

func payTax2(salary int) {
	if salary < 150000 {
		fmt.Println(errors.New("error: el salario ingresado no alcanza el mínimo imponible"))
		return
	}
	fmt.Println("Debe pagar impuesto.")
}

//Ejercicio 3
//Repite el proceso anterior, pero ahora implementando 'fmt.Errorf()',
//para que el mensaje de error reciba por parámetro el valor de 'salary' indicando que no alcanza
//el mínimo imponible (el mensaje mostrado por consola deberá decir: 'error: el mínimo imponible
//es de 150.000 y el salario ingresado es de: [salary]', siendo [salary] el valor de tipo int pasado por parámetro).

func payTax3(salary int) {
	if salary < 150000 {
		errorSalary := fmt.Errorf("error: el mínimo imponible es de 150.000 y el salario ingresado es de: %v", salary)
		fmt.Println(errorSalary)
		return
	}
	fmt.Println("Debe pagar impuesto.")
}

func main() {

	salary = 160000
	fmt.Println("Resultado ejercicio 1")
	msg, err := payTax1(salary)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(msg)

	fmt.Println("Resultado ejercicio 2")
	payTax2(salary)

	fmt.Println("Resultado ejercicio 3")
	payTax3(salary)
}
