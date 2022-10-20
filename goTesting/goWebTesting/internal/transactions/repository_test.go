package transactions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {

	//arrange

	database := []Transaction{
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

	studStorage := StubDB{
		dataStub: database,
	}
	repo := NewRepository(&studStorage)

	//act

	result, err := repo.GetAll()

	//result = append(result, Transaction{}) // Agrega una Transaction vacia al resultado de GetAll para que no coincida
	// con la lista de Transactions de la db y falle el Test

	//assert
	assert.Nil(t, err)
	assert.Equal(t, database, result)
}
