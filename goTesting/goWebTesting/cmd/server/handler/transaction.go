package handler

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/juanimeli/backpack-bcgow6-juan-guglielmone/goTesting/goWebTesting/internal/transactions"
	"github.com/juanimeli/backpack-bcgow6-juan-guglielmone/goTesting/goWebTesting/pkg/web"
)

type request struct {
	Codigo   string  `json:"cod"`      // binding:"required"`
	Moneda   string  `json:"currency"` // binding:"required"`
	Monto    float64 `json:"amount"`   // binding:"required"`
	Emisor   string  `json:"sender"`   //binding:"required"`
	Receptor string  `json:"receiver"` // binding:"required"`
	Fecha    string  `json:"date"`     //binding:"required"`
}

type Transaction struct {
	service transactions.Service
}

func NewTransaction(t transactions.Service) *Transaction {
	return &Transaction{
		service: t,
	}
}

// StoreNewTransaction godoc
// @Summary List new Transaction
// @Tags Transactions
// @Description post new Transaction and save into de db json file
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param transaction body request true "Transaction to store"
// @Success 200 {object} web.Response "Transaction Stored"
// @Failure 401 {object} web.Response "Unauthorized"
// @Failure 400 {object} web.Response "Invalid parameter"
// @Router /transactions [post]
func (c *Transaction) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "invalid token"))
			return
		}

		var r request
		if err := ctx.ShouldBindJSON(&r); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("field %s is required", err.Error())))
			return
		}

		t, err := c.service.Store(r.Codigo, r.Moneda, r.Monto, r.Emisor, r.Receptor, r.Fecha)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		ctx.JSON(201, web.NewResponse(201, t, ""))
	}

}

// ListTransactions godoc
// @Summary List existing transactions
// @Tags Transactions
// @Description get transactions from json db file
// @Produce json
// @Success 200 {object} web.Response "List Transactions"
// @Failure 404 {object} web.Response "Transactions not found"
// @Router /transactions [get]
func (c *Transaction) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
		}
		ctx.JSON(200, web.NewResponse(200, t, ""))
	}

}

// UpdateTransaction godoc
// @Summary Update Transaction
// @Tags Transactions
// @Description update all the parameters of an existing Transaction of the db json file
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param transaction body request true "New paramaters to update the existing Transaction"
// @Success 200 {object} web.Response "Transaction Updated"
// @Failure 401 {object} web.Response "Unauthorized"
// @Failure 400 {object} web.Response "Invalid parameter"
// @Failure 404 {object} web.Response "Transaction not found"
// @Router /transactions/{ID} [put]
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

// DeleteTransaction godoc
// @Summary Delete Transaction
// @Tags Transactions
// @Description Delete an existing Transaction of the db json file
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} web.Response "Transaction Deleted"
// @Failure 401 {object} web.Response "Unauthorized"
// @Failure 400 {object} web.Response "Invalid parameter"
// @Failure 404 {object} web.Response "Transaction not found"
// @Router /transactions/{ID} [delete]
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
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, fmt.Sprintf("transaction with id %d has been deleted", id), ""))

	}
}

// UpdateCodnAmount Transaction godoc
// @Summary Partial Update on a Transaction
// @Tags Transactions
// @Description update cod and amount parameters of an existing Transaction of the db json file
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param transaction body request true "New paramaters to update the existing Transaction"
// @Success 200 {object} web.Response "cod and amount Transaction Updated"
// @Failure 401 {object} web.Response "Unauthorized"
// @Failure 400 {object} web.Response "Invalid parameter"
// @Failure 404 {object} web.Response "Transaction not found"
// @Router /transactions/{ID} [patch]
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
