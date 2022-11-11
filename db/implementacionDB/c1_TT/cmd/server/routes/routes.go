package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/juanimeli/backpack-bcgow6-juan-guglielmone/db/implementaciondb/c1_tt/cmd/server/handler"
	"github.com/juanimeli/backpack-bcgow6-juan-guglielmone/db/implementaciondb/c1_tt/internal/product"
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
	return &router{r: r, db: db}
}

func (r *router) MapRoutes() {
	r.setGroup()
	r.buildSellerRoutes()
}

func (r *router) setGroup() {
	r.rg = r.r.Group("/api/v1")
}

func (r *router) buildSellerRoutes() {
	repo := product.NewRepo(r.db)
	service := product.NewService(repo)
	handler := handler.NewProduct(service)

	r.rg.GET("/products", handler.GetAll())
	r.rg.GET("/products/:id", handler.Get())
	r.rg.GET("/products/complete/:id", handler.GetFullData())
	r.rg.POST("/products", handler.Create())
	r.rg.DELETE("/products/:id", handler.Delete())
	r.rg.PATCH("/products/:id", handler.Update())
}
