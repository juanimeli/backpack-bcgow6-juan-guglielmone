package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/juanimeli/backpack-bcgow6-juan-guglielmone/goTesting/goWebTesting/cmd/server/handler"
	"github.com/juanimeli/backpack-bcgow6-juan-guglielmone/goTesting/goWebTesting/internal/transactions"
)

func createServer(MockStorage transactions.MockDB) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	repo := transactions.NewRepository(&MockStorage)
	service := transactions.NewService(repo)
	t := handler.NewTransaction(service)

	r := gin.Default()

	r.GET("/transactions/", t.GetAll())
	r.POST("/transactions/", t.Store())

	return r
}

func createRequestTest(method, url, body string) (*http.Request, *httptest.ResponseRecorder) {

	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	//req.Header.Add("token", "12345")

	return req, httptest.NewRecorder()
}

func TestGetAllTransactions(t *testing.T) {

	//arrange
	database := []transactions.Transaction{
		{
			ID:       123,
			Codigo:   "asd",
			Moneda:   "USD",
			Monto:    40.00,
			Emisor:   "Juan",
			Receptor: "Pedro",
			Fecha:    "23/10/2022",
		},
		{
			ID:       124,
			Codigo:   "asda",
			Moneda:   "USD",
			Monto:    44.00,
			Emisor:   "Pedro",
			Receptor: "Juan",
			Fecha:    "24/10/2022",
		},
	}

	MockStorage := transactions.MockDB{
		DataMock: []transactions.Transaction{
			{
				ID:       123,
				Codigo:   "asd",
				Moneda:   "USD",
				Monto:    40.00,
				Emisor:   "Juan",
				Receptor: "Pedro",
				Fecha:    "23/10/2022",
			},
			{
				ID:       124,
				Codigo:   "asda",
				Moneda:   "USD",
				Monto:    44.00,
				Emisor:   "Pedro",
				Receptor: "Juan",
				Fecha:    "24/10/2022",
			},
		},
	}

	var resp []transactions.Transaction
	r := createServer(MockStorage)
	req, rr := createRequestTest(http.MethodGet, "/transactions/", "")

	//act
	r.ServeHTTP(rr, req)
	err := json.Unmarshal(rr.Body.Bytes(), &resp)

	//assert

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, len(database), len(resp))

}

func TestSaveTransaction(t *testing.T) {

	//arrange
	mockStorage := transactions.MockDB{
		DataMock: []transactions.Transaction{
			{
				ID:       123,
				Codigo:   "asd",
				Moneda:   "USD",
				Monto:    40.00,
				Emisor:   "Juan",
				Receptor: "Pedro",
				Fecha:    "23/10/2022",
			},
			{
				ID:       124,
				Codigo:   "asda",
				Moneda:   "USD",
				Monto:    44.00,
				Emisor:   "Pedro",
				Receptor: "Juan",
				Fecha:    "24/10/2022",
			},
		},
	}
	var resp transactions.Transaction
	r := createServer(mockStorage)
	req, rr := createRequestTest(http.MethodPost, "/transactions/", `{
		cod: "NEW", currency: "NEW", amount: 44.00, sender: "Pedro", receiver: "Juan", date: "24/10/2022",
	}`)
	//act
	r.ServeHTTP(rr, req)

	//assert
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, rr.Code)

}
