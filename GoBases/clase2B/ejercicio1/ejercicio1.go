package main

import "fmt"

//Una universidad necesita registrar a los/as estudiantes y generar una funcionalidad para imprimir el
//detalle de los datos de cada uno de ellos/as, de la siguiente manera:
//  Nombre: [Nombre del alumno]
//  Apellido: [Apellido del alumno]
//  DNI: [DNI del alumno]
//  Fecha: [Fecha ingreso alumno]
//Los valores que están en corchetes deben ser reemplazados por los datos brindados por los alumnos/as.
//Para ello es necesario generar una estructura Alumnos con las variables Nombre, Apellido, DNI, Fecha
//y que tenga un método detalle

type student struct {
	nombre   string
	apellido string
	dni      string
	fecha    string
}

func (s *student) setStudent(nombre, apellido, dni, fecha string) {
	s.nombre = nombre
	s.apellido = apellido
	s.dni = dni
	s.fecha = fecha
	fmt.Printf("Datos del estudiante:\nNombre: %v\nApellido: %v\nDni: %v\nFecha: %v\n", s.nombre, s.apellido, s.dni, s.fecha)
}

func main() {
	student1 := student{}
	student1.setStudent("nani", "martin", "112233", "10/10/22")

	student2 := student{}
	student2.setStudent("juli", "rios", "998877", "21/08/22")
}
