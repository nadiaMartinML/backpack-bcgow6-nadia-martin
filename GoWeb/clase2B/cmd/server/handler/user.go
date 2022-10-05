package handler

import (
	//"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/nadiaMartinML/backpack-bcgow6-nadia-martin/GoWeb/clase2B/internal/users"
)

type request struct {
	Id         int     `json:"Id" binding:"required"`
	Name       string  `form:"Name" json:"Name" validate:"required"`
	LastName   string  `form:"LastName" json:"LastName" validate:"required"`
	Email      string  `form:"Email" json:"Email" validate:"required"`
	Age        int     `form:"Age" json:"Age" validate:"required"`
	Height     float64 `form:"Height" json:"Height" validate:"required"`
	Active     bool    `form:"Active" json:"Active" validate:"required"`
	DateCreate string  `form:"DateCreate" json:"DateCreate" validate:"required"`
}

type User struct {
	service users.Service
}

func NewUser(u users.Service) *User {
	return &User{
		service: u,
	}
}

func (r *request) validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

func (u *User) GetAllUsers(c *gin.Context) {
	token := c.GetHeader("token")
	if token != "password" || token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "no tiene permisos para realizar la petición solicitada"})
		return
	}

	users, err := u.service.GetAllUsers()
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	if len(users) <= 0 {
		c.JSON(http.StatusNoContent, gin.H{"error": "No hay usuarios cargados"})
		return
	} else {
		c.JSON(200, users)
	}
}

func (u *User) SaveUser(c *gin.Context) {

	token := c.GetHeader("token")
	if token != "password" || token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "no tiene permisos para realizar la petición solicitada"})
		return
	}

	var req request

	c.ShouldBindJSON(&req)
	if err := req.validate(); err != nil {
		valError := err.(validator.ValidationErrors)
		msg := []string{}
		for _, e := range valError {
			msg = append(msg, "el campo "+e.Field()+" es requerido")
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": msg})
		return
	}

	userSave, err := u.service.SaveUser(req.Name, req.LastName, req.Email, req.Age, req.Height, req.Active, req.DateCreate)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, userSave)
}
