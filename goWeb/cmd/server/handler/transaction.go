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

	"github.com/gin-gonic/gin"
	"github.com/juanimeli/backpack-bcgow6-juan-guglielmone/goWeb/internal/transactions"
)

const (
	filePath = "./transactions.json"
	Token    = "12345"
)

type request struct {
	Codigo   string  `json:"cod" binding:"required"`
	Moneda   string  `json:"currency" binding:"required"`
	Monto    float64 `json:"amount" binding:"required"`
	Emisor   string  `json:"sender" binding:"required"`
	Receptor string  `json:"receiver" binding:"required"`
	Fecha    string  `json:"date" binding:"required"`
}

type Transaction struct {
	service transactions.Service
}

func NewTransaction(t transactions.Service) *Transaction {
	return &Transaction{
		service: t,
	}
}

func (c *Transaction) AddTransaction() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != Token {
			ctx.JSON(401, gin.H{"error": "invalid token"})
			return
		}

		var r request
		if err := ctx.ShouldBindJSON(&r); err != nil {
			ctx.JSON(404, gin.H{"error": fmt.Errorf("field %s is required", err.Error())})
			return
		}

		//t, err := c.service.AddTransaction(r.Codigo, r.Moneda, r.Monto, r.Emisor, r.Receptor, r.Fecha)

		ctx.JSON(200, gin.H{"transaction added": r})
	}

}

func (c *Transaction) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
		}
		ctx.JSON(200, gin.H{"content": t})
	}

	/*transactions, err := ReadJson(filePath)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}*/

}
