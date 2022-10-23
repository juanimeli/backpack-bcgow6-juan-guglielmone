package transactions

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Para probar la integracion del repositorio con el servicio, y testear particularmente
// el metood GetAll() (si esta bien integrado), se repite el procedimiento del test unitario
// del repositorio nada mas que se llama a GetAll desde servicio. Es decir estamos probando justamente
// como funciona repositorio con servicio integrados. Para eso generamos primero el Mock que imita
// a la base de datos y luego validamos que el repositorio creado con el Mock y el servicio instanciasdo
// con el repo devuelva en el metodo GetAll() de service lo que se espera.
func TestGetAllServiceIntegration(t *testing.T) {
	//arrange

	database := []Transaction{
		{
			ID:       1,
			Codigo:   "Primer",
			Moneda:   "USD",
			Monto:    40.00,
			Emisor:   "Juan",
			Receptor: "Pedro",
			Fecha:    "23/10/2022",
		},
		{
			ID:       2,
			Codigo:   "Segunda",
			Moneda:   "USD",
			Monto:    44.00,
			Emisor:   "Pedro",
			Receptor: "Juan",
			Fecha:    "24/10/2022",
		},
	}

	MockStorage := MockDB{
		dataMock: database,
	}
	repo := NewRepository(&MockStorage)
	service := NewService(repo)

	//act

	result, err := service.GetAll()

	//result = append(result, Transaction{})
	// Agrega una Transaction vacia al resultado de GetAll para que no coincida
	// con la lista de Transactions de la db y falle el Test

	//assert
	assert.Nil(t, err)
	assert.Equal(t, database, result)
}

func TestGetAllServiceIntegrationFail(t *testing.T) {
	//arrange
	errorEsperado := errors.New("Soy un error de lectura >:(")
	MockStorage := MockDB{
		dataMock:   nil,
		errOnRead:  errors.New("Soy un error de lectura >:("),
		errOnWrite: nil,
	}
	repo := NewRepository(&MockStorage)
	service := NewService(repo)

	//act

	result, err := service.GetAll()

	//assert
	assert.Equal(t, errorEsperado, err)
	assert.Empty(t, result)
}
func TestStoreServiceIntegration(t *testing.T) {
	//arrange

	database := []Transaction{
		{
			ID:       1,
			Codigo:   "asd",
			Moneda:   "USD",
			Monto:    40.00,
			Emisor:   "Juan",
			Receptor: "Pedro",
			Fecha:    "23/10/2022",
		},
		{
			ID:       2,
			Codigo:   "asda",
			Moneda:   "USD",
			Monto:    44.00,
			Emisor:   "Pedro",
			Receptor: "Juan",
			Fecha:    "24/10/2022",
		},
	}

	esperado := Transaction{
		ID:       3,
		Codigo:   "NEW",
		Moneda:   "UYU",
		Monto:    20.00,
		Emisor:   "Juan",
		Receptor: "Martin",
		Fecha:    "22/10",
	}

	MockStorage := MockDB{
		dataMock: database,
	}
	repo := NewRepository(&MockStorage)
	service := NewService(repo)

	//act

	result, err := service.Store("NEW", "UYU", 20.00, "Juan", "Martin", "22/10")

	//assert
	assert.Nil(t, err)
	assert.Equal(t, esperado, result)
}

func TestStoreServiceIntegrationFail(t *testing.T) {
	// arrange
	errorEsperado := errors.New("Soy un error de escritura")
	mockStorage := MockDB{
		dataMock:   nil,
		errOnRead:  nil,
		errOnWrite: errors.New("Soy un error de escritura"),
	}
	repo := NewRepository(&mockStorage)
	service := NewService(repo)

	newTrans := Transaction{
		ID:       3,
		Codigo:   "NEW",
		Moneda:   "UYU",
		Monto:    20.00,
		Emisor:   "Juan",
		Receptor: "Martin",
		Fecha:    "22/10",
	}

	// act

	result, err := service.Store(newTrans.Codigo, newTrans.Moneda, newTrans.Monto, newTrans.Emisor, newTrans.Receptor, newTrans.Fecha)

	// assert

	assert.Equal(t, errorEsperado, err)
	assert.Empty(t, result)

}
