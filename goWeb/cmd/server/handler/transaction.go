package handler

/*
El paquete handler con el controlador de la entidad seleccionada.
Se debe generar la estructura request
Se debe generar la estructura del controlador que tenga como campo el servicio
Se debe generar la función que retorne el controlador
Se deben generar todos los métodos correspondientes a los endpoints
*/

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/juanimeli/backpack-bcgow6-juan-guglielmone/goWeb/internal/transactions"
	"github.com/juanimeli/backpack-bcgow6-juan-guglielmone/goWeb/pkg/web"
)

type request struct { //la estructura request es igual
	//a la de la entidad sin el ID
	Codigo   string  `json:"cod"`      // binding:"required"`
	Moneda   string  `json:"currency"` // binding:"required"`
	Monto    float64 `json:"amount"`   // binding:"required"`
	Emisor   string  `json:"sender"`   //binding:"required"`
	Receptor string  `json:"receiver"` // binding:"required"`
	Fecha    string  `json:"date"`     //binding:"required"`
}

type Transaction struct {
	service transactions.Service //la estructura dela entidad tiene como atributo
	//la estructura del servicio
	//y la estructura del servicio tiene
	//como atributo la estructura del repositorio
	//la estructura del repo esta vacia pero contendria
	// la base de datos.
}

func NewTransaction(t transactions.Service) *Transaction {
	return &Transaction{
		service: t,
	}
}

func (c *Transaction) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "invalid token"))
			return
		}

		var r request
		if err := ctx.ShouldBindJSON(&r); err != nil {
			ctx.JSON(404, web.NewResponse(401, nil, fmt.Sprintf("field %s is required", err.Error())))
			return
		}

		t, err := c.service.Store(r.Codigo, r.Moneda, r.Monto, r.Emisor, r.Receptor, r.Fecha)
		if err != nil {
			ctx.JSON(401, web.NewResponse(401, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, t, ""))
	}

}

func (c *Transaction) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
		}
		ctx.JSON(200, web.NewResponse(200, t, ""))
	}

}
func (c *Transaction) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "invalid token"))
			return
		}
		id, err := strconv.ParseInt(ctx.Param("ID"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "invalid ID"))
			return
		}
		var r request

		if err := ctx.ShouldBindJSON(&r); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		if r.Codigo == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "cod is required"))
			return
		}
		if r.Moneda == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "currency is required"))
			return
		}
		if r.Monto == 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "amount is required"))
			return
		}
		if r.Emisor == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "sender is required"))
			return
		}
		if r.Receptor == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "receiver is required"))
			return
		}
		if r.Fecha == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "date is required"))
			return
		}
		t, err := c.service.Update(int(id), r.Codigo, r.Moneda, r.Monto, r.Emisor, r.Receptor, r.Fecha)
		if err != nil {
			ctx.JSON(404, web.NewResponse(400, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, t, ""))
	}
}

func (c *Transaction) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "invalid token"))
			return
		}
		id, err := strconv.ParseInt(ctx.Param("ID"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "invalid ID"))
			return
		}
		err = c.service.Delete(int(id))
		if err != nil {
			ctx.JSON(401, web.NewResponse(401, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, fmt.Sprintf("transaction with id %d has been deleted", id), ""))

	}
}
func (c *Transaction) UpdateCodnAmount() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "invalid token"))
			return
		}
		id, err := strconv.ParseInt(ctx.Param("ID"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "invalid ID"))
			return
		}
		var r request
		if err := ctx.ShouldBindJSON(&r); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		if r.Codigo == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "transaction cod is required"))
			return
		}
		if r.Monto == 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "transaction amount is required"))
		}

		t, err := c.service.UpdateCodnAmount(int(id), r.Codigo, r.Monto)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, t, ""))
	}
}

/*transactions, err := ReadJson(filePath)
if err != nil {
	ctx.JSON(400, err.Error())
	return
}*/
/*
	p, err := c.service.Update(int(id), req.Name, req.Type, req.Count, req.Price)
	if err != nil {
		ctx.JSON(404, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, p)
}*/
