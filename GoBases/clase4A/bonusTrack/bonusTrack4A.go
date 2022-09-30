package main

import (
	"errors"
	"fmt"
)

func calculateSalary(hours, priceHour float64) (salary float64, errHours error) {
	salary = hours * priceHour
	if salary >= 150000 {
		salary = salary * 0.1
	}
	if priceHour < 80 {
		errHours = errors.New("error: el trabajador no puede haber trabajado menos de 80hs mensuales")
	}
	return salary, nil
}

func calculateBonus(bestSalary, monthsWork float64) (halfBonus float64, errBonus error) {
	bonus := bestSalary / 12 * monthsWork
	if bestSalary < 0 || monthsWork < 0 {
		errBonus = fmt.Errorf("error: número negativo")
	}
	halfBonus = bonus / 2
	return halfBonus, nil
}

func main() {
	salary, errHours := calculateSalary(100, 50)
	halfBonus, errBonus := calculateBonus(6000, 8)

	if errHours != nil {
		fmt.Println(errHours)
	}

	if errBonus != nil {
		fmt.Println(errBonus)
	}
	fmt.Println("El salario mensual es de: ", salary)
	fmt.Println("El medio aguinaldo es de: ", halfBonus)

}
