package main

import "fmt"

//Una empresa de redes sociales requiere implementar una estructura usuario con funciones que vayan
//agregando información a la estructura. Para optimizar y ahorrar memoria requieren que la estructura
//de usuarios ocupe el mismo lugar en memoria para el main del programa y para las funciones.
//La estructura debe tener los campos: Nombre, Apellido, Edad, Correo y Contraseña
//Y deben implementarse las funciones:
//Cambiar nombre: me permite cambiar el nombre y apellido.
//Cambiar edad: me permite cambiar la edad.
//Cambiar correo: me permite cambiar el correo.
//Cambiar contraseña: me permite cambiar la contraseña.

type user struct {
	name     string
	lastName string
	age      int
	email    string
	password string
}

func (u *user) changeName(name *string, lastName *string) {
	u.name = *name
	u.lastName = *lastName
}

func (u *user) changeAge(age *int) {
	u.age = *age
}

func (u *user) changeEmail(email *string) {
	u.email = *email
}

func (u *user) changePassword(password *string) {
	u.password = *password
}

func main() {
	u1 := user{
		name:     "nani",
		lastName: "martin",
		age:      31,
		email:    "test@test.com",
		password: "pass",
	}
	fmt.Printf("Nombre del usuario es: %v %v, su edad es: %v, su mail es: %v y su pass es: %v", u1.name, u1.lastName, u1.age, u1.email, u1.password)
	fmt.Print("\n")

	name := "nadia"
	lastName := "montesi"
	age := 33
	email := "nmm@mail.com"
	password := "password"

	u1.changeName(&name, &lastName)
	u1.changeAge(&age)
	u1.changeEmail(&email)
	u1.changePassword(&password)

	fmt.Printf("Nombre del usuario es: %v %v, su edad es: %v, su mail es: %v y su pass es: %v", u1.name, u1.lastName, u1.age, u1.email, u1.password)
	fmt.Print("\n")

}
