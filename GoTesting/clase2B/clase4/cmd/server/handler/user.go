package handler

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/nadiaMartinML/backpack-bcgow6-nadia-martin/GoTesting/clase2B/clase4/internal/users"
	"github.com/nadiaMartinML/backpack-bcgow6-nadia-martin/GoTesting/clase2B/clase4/pkg/web"
)

type request struct {
	Id         int     `json:"Id"`
	Name       string  `form:"Name" json:"Name" validate:"required"`
	LastName   string  `form:"LastName" json:"LastName" validate:"required"`
	Email      string  `form:"Email" json:"Email" validate:"required"`
	Age        int     `form:"Age" json:"Age" validate:"required"`
	Height     float64 `form:"Height" json:"Height" validate:"required"`
	Active     *bool   `form:"Active" json:"Active" validate:"required"`
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

// ListUsers
// @Summary List users
// @Tags Users
// @Description get users
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Failure 401 {object} web.Response
// @Failure 400 {object} web.Response
// @Router /users [get]
func (u *User) GetAllUsers(c *gin.Context) {
	token := c.GetHeader("token")
	if token != os.Getenv("TOKEN") {
		c.JSON(http.StatusUnauthorized, web.NewResponse(401, nil, "No tiene los permisos para realizar la petición solicitada"))
		return
	}

	users, err := u.service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "No se pudo acceder a la información"))
		return
	}

	if len(users) <= 0 {
		c.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "No hay información para mostrar"))
		return
	} else {
		c.JSON(http.StatusOK, web.NewResponse(200, users, ""))
	}
}

// Create Users
// @Summary save users
// @Tags Users
// @Description save users
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param user body request true "User to store"
// @Success 200 {object} web.Response
// @Failure 401 {object} web.Response
// @Failure 404 {object} web.Response
// @Router /users [post]
func (u *User) SaveUser(c *gin.Context) {

	token := c.GetHeader("token")
	if token != os.Getenv("TOKEN") {
		c.JSON(http.StatusUnauthorized, web.NewResponse(401, nil, "No tiene los permisos para realizar la petición solicitada"))
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
		c.JSON(http.StatusNotFound, web.NewResponse(404, nil, fmt.Sprint(msg)))
		return
	}

	userSave, err := u.service.SaveUser(req.Name, req.LastName, req.Email, req.Age, req.Height, *req.Active, req.DateCreate)
	if err != nil {
		c.JSON(http.StatusNotFound, web.NewResponse(404, nil, "Error al cargar un nuevo usuario"))
		return
	}

	c.JSON(http.StatusOK, web.NewResponse(200, userSave, ""))
}

// Update User
// @Summary update user
// @Tags Users
// @Description update user
// @Accept json
// @Produce json
// @Param id path int true "User id"
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Failure 401 {object} web.Response
// @Failure 400 {object} web.Response
// @Router /users/{id} [put]
func (u *User) UpdateUser(c *gin.Context) {
	token := c.GetHeader("token")
	if token != os.Getenv("TOKEN") {
		c.JSON(http.StatusUnauthorized, web.NewResponse(401, nil, "No tiene los permisos para realizar la petición solicitada"))
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "Id inválido"))
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
		c.JSON(http.StatusBadRequest, web.NewResponse(400, nil, fmt.Sprint(msg)))
		return
	}

	userUpdate, err := u.service.UpdateUser(int(id), req.Name, req.LastName, req.Email, req.Age, req.Height, *req.Active, req.DateCreate)
	if err != nil {
		c.JSON(http.StatusBadRequest, web.NewResponse(400, nil, err.Error()))
		return
	}
	c.JSON(http.StatusOK, web.NewResponse(200, userUpdate, ""))
}

// Delete User
// @Summary delete user
// @Tags Users
// @Description delete user
// @Param id path int true "User id"
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Failure 401 {object} web.Response
// @Failure 400 {object} web.Response
// @Failure 404 {object} web.Response
// @Router /users/{id} [delete]
func (u *User) DeleteUser(c *gin.Context) {
	token := c.GetHeader("token")
	if token != os.Getenv("TOKEN") {
		c.JSON(http.StatusUnauthorized, web.NewResponse(401, nil, "No tiene los permisos para realizar la petición solicitada"))
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "Id inválido"))
		return
	}

	err = u.service.DeleteUser(int(id))
	if err != nil {
		c.JSON(http.StatusNotFound, web.NewResponse(404, nil, err.Error()))
		return
	}

	c.JSON(http.StatusOK, web.NewResponse(200, fmt.Sprintf("El usuario %d ha sido eliminado correctamente", id), ""))
}

// UpdatePatch User
// @Summary update partial user
// @Tags Users
// @Description update partial user
// @Accept json
// @Produce json
// @Param id path int true "User id"
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Failure 401 {object} web.Response
// @Failure 400 {object} web.Response
// @Failure 404 {object} web.Response
// @Router /users/{id} [patch]
func (u *User) UpdatePatch(c *gin.Context) {
	token := c.GetHeader("token")
	if token != os.Getenv("TOKEN") {
		c.JSON(http.StatusUnauthorized, web.NewResponse(401, nil, "No tiene los permisos para realizar la petición solicitada"))
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(404, web.NewResponse(404, nil, "El id es inválido"))
		return
	}

	var req request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(404, web.NewResponse(404, nil, err.Error()))
		return
	}
	if req.LastName == "" {
		c.JSON(404, web.NewResponse(404, nil, "El campo Apellido no puede estar vacío"))
		return
	}
	if req.Age <= 0 {
		c.JSON(404, web.NewResponse(404, nil, "El campo Edad no puede ser menor a 1"))
		return
	}

	userUpdate, err := u.service.UpdatePatch(int(id), req.LastName, req.Age)
	if err != nil {
		c.JSON(http.StatusBadRequest, web.NewResponse(400, nil, err.Error()))
		return
	}
	c.JSON(http.StatusOK, web.NewResponse(200, userUpdate, ""))
}
