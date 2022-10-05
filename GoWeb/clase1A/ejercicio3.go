package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id         int     `json:"Id"`
	Name       string  `form:"Name" json:"Name"`
	LastName   string  `form:"LastName" json:"LastName"`
	Email      string  `form:"Email" json:"Email"`
	Age        int     `form:"Age" json:"Age"`
	Height     float64 `form:"Height" json:"Height"`
	Active     bool    `form:"Active" json:"Active"`
	DateCreate string  `form:"DateCreate" json:"DateCreate"`
}

var Users []User

func getAllUsers(c *gin.Context) {
	c.JSON(200, Users)
}

// ejercicio 1 clase TT
func userFilter(c *gin.Context) {
	allUsers := Users
	var filtered []User

	name := c.Query("Name")
	lastName := c.Query("LastName")
	email := c.Query("Email")
	age := c.Query("Age")
	height := c.Query("Height")
	active := c.Query("Active")
	dateCreate := c.Query("DateCreate")

	ageKey, _ := strconv.Atoi(age)
	heightKey, _ := strconv.ParseFloat(email, 64)
	activeKey, _ := strconv.ParseBool(active)

	for _, user := range allUsers {
		if (name != "" && name == user.Name) || (lastName != "" && lastName == user.LastName) || (email != "" && email == user.Email) || (age != "" && ageKey == user.Age) || (height != "" && heightKey == user.Height) || (active != "" && activeKey == user.Active) || (dateCreate != "" && dateCreate == user.DateCreate) {
			filtered = append(filtered, user)
			return
		}
		fmt.Println(allUsers)
		fmt.Println(filtered)
	}
	c.JSON(200, filtered)
}

func getOneUser(c *gin.Context) {
	id, errId := strconv.Atoi(c.Param("id"))
	if errId != nil {
		return
	}

	for _, u := range Users {
		if u.Id == id {
			c.JSON(200, u)
			return
		}
	}
	c.JSON(404, "Información del usuario ¡No existe!")
	return
}

func main() {

	data, err := os.ReadFile("../users.json")
	if err != nil {
		panic("Hubo un error en la lectura del archivo...")
	}

	if err := json.Unmarshal(data, &Users); err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	rg := router.Group("users")
	rg.GET("/GetAll", getAllUsers)
	rg.GET("/:id", getOneUser)
	rg.GET("/filter", userFilter)

	router.Run()
}
