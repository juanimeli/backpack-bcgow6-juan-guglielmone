package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/juanimeli/backpack-bcgow6-juan-guglielmone/goTesting/goWebTesting/cmd/server/handler"
	"github.com/juanimeli/backpack-bcgow6-juan-guglielmone/goTesting/goWebTesting/internal/transactions"
	"github.com/juanimeli/backpack-bcgow6-juan-guglielmone/goTesting/goWebTesting/pkg/store"
	"github.com/juanimeli/backpack-bcgow6-juan-guglielmone/goTesting/goWebTesting/pkg/web"
)

var s = createServer()

func createServer() *gin.Engine {

	err := os.Setenv("TOKEN", "12345")
	if err != nil {
		panic(err)
	}
	db := store.New(store.FileType, "../transactionstest.json")
	repo := transactions.NewRepository(db)
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
	req.Header.Add("token", os.Getenv("TOKEN"))

	return req, httptest.NewRecorder()
}

func TestGetAllTransactions(t *testing.T) {

	//arrange

	req, rr := createRequestTest(http.MethodGet, "/transactions/", "")

	//act
	s.ServeHTTP(rr, req)
	objRes := &web.Response{}
	assert.Equal(t, http.StatusOK, rr.Code)
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)

	data := reflect.ValueOf(objRes.Data).Len()
	assert.Nil(t, err)
	assert.True(t, data > 0)

	//assert

}

func TestSaveTransaction(t *testing.T) {

	//arrange

	var resp transactions.Transaction
	r := createServer()
	req, rr := createRequestTest(http.MethodPost, "/transactions/", `{"cod": "NEW", "currency": "NEW", "amount": 44.00, "sender": "Pedro", "receiver": "Juan", "date": "24/10/2022"}`)
	//act
	r.ServeHTTP(rr, req)
	json.Unmarshal(rr.Body.Bytes(), &resp)
	//assert

	assert.Equal(t, http.StatusCreated, rr.Code)

}
