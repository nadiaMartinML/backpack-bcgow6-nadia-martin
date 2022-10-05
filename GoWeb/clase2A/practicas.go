package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type User struct {
	Id         int     `json:"Id" binding:"required"`
	Name       string  `form:"Name" json:"Name" validate:"required"`
	LastName   string  `form:"LastName" json:"LastName" validate:"required"`
	Email      string  `form:"Email" json:"Email" validate:"required"`
	Age        int     `form:"Age" json:"Age" validate:"required"`
	Height     float64 `form:"Height" json:"Height" validate:"required"`
	Active     bool    `form:"Active" json:"Active" validate:"required"`
	DateCreate string  `form:"DateCreate" json:"DateCreate" validate:"required"`
}

var Users []User

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
	rg.POST("/", saveUser)

	router.Run()
}

func (u *User) validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

func saveUser(c *gin.Context) {
	var user User

	token := c.GetHeader("token")
	if token == "" || token != "password" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "no tiene permisos para realizar la petición solicitada"})
		return
	}

	user.Id = len(Users) + 1

	c.ShouldBindJSON(&user)
	if err := user.validate(); err != nil {
		valError := err.(validator.ValidationErrors)
		msg := []string{}
		for _, e := range valError {
			msg = append(msg, "el campo "+e.Field()+" es requerido")
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": msg})
		return
	}

	Users = append(Users, user)
	c.JSON(http.StatusOK, user)
}
