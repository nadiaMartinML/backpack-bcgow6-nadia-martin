package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/nadiaMartinML/backpack-bcgow6-nadia-martin/Storage/clase1/cmd/server/handler"
	"github.com/nadiaMartinML/backpack-bcgow6-nadia-martin/Storage/clase1/product"
)

type Router interface {
	MapRoutes()
}

type router struct {
	r  *gin.Engine
	rg *gin.RouterGroup
	db *sql.DB
}

func NewRouter(r *gin.Engine, db *sql.DB) Router {
	return &router{
		r:  r,
		db: db,
	}
}

func (r *router) MapRoutes() {
	r.setGroup()
	r.buildSellerRoutes()
}

func (r *router) setGroup() {
	r.rg = r.r.Group("/api/v1")
}

func (r *router) buildSellerRoutes() {
	repo := product.NewRepository(r.db)
	service := product.NewService(repo)
	handler := handler.NewProduct(service)
	r.rg.GET("products/:id", handler.GetOne)
	r.rg.GET("products", handler.GetAll)
	r.rg.POST("products", handler.Store)
	r.rg.PATCH("products/:id", handler.Update)
	r.rg.DELETE("products/:id", handler.Delete)
	r.rg.GET("productsWithWarehouse/:id", handler.GetFullData)
}
