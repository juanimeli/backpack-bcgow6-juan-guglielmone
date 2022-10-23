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

func TestUpdateServiceIntegration(t *testing.T) {
	database := []Transaction{
		{
			ID:       1,
			Codigo:   "BEFORE UPDATE",
			Moneda:   "USD",
			Monto:    10.00,
			Emisor:   "Juan",
			Receptor: "Pedro",
			Fecha:    "23/10/2022",
		},
	}
	//arrange
	mockStorage := MockDB{
		dataMock:   database,
		errOnRead:  nil,
		errOnWrite: nil,
		readCheck:  false,
	}
	repo := NewRepository(&mockStorage)
	service := NewService(repo)

	expected := Transaction{
		ID:       1,
		Codigo:   "AFTER UPDATE",
		Moneda:   "USD AFTER",
		Monto:    20.00,
		Emisor:   "Juan",
		Receptor: "Pedro",
		Fecha:    "24/10/2022",
	}

	//act

	result, err := service.Update(expected.ID, expected.Codigo, expected.Moneda, expected.Monto, expected.Emisor, expected.Receptor, expected.Fecha)

	//assert
	assert.Nil(t, err)
	assert.True(t, mockStorage.readCheck)
	assert.Equal(t, expected, result)

}

func TestUpdateServiceIntegrationReadFail(t *testing.T) {

	//arrange
	errorEsperado := errors.New("Soy un error de Read")
	mockStorage := MockDB{
		dataMock:   nil,
		errOnRead:  errors.New("Soy un error de Read"),
		errOnWrite: nil,
		readCheck:  false,
	}
	repo := NewRepository(&mockStorage)
	service := NewService(repo)

	tUpdated := Transaction{
		ID:       1,
		Codigo:   "AFTER UPDATE",
		Moneda:   "USD AFTER",
		Monto:    20.00,
		Emisor:   "Juan",
		Receptor: "Pedro",
		Fecha:    "24/10/2022",
	}

	//act

	result, err := service.Update(tUpdated.ID, tUpdated.Codigo, tUpdated.Moneda,
		tUpdated.Monto, tUpdated.Emisor, tUpdated.Receptor, tUpdated.Fecha)

	//assert

	assert.False(t, mockStorage.readCheck)
	assert.Equal(t, errorEsperado, err)
	assert.Empty(t, result)

}

func TestUpdateServiceIntegrationWriteFail(t *testing.T) {

	//arrange
	database := []Transaction{
		{
			ID:       1,
			Codigo:   "BEFORE UPDATE",
			Moneda:   "USD",
			Monto:    10.00,
			Emisor:   "Juan",
			Receptor: "Pedro",
			Fecha:    "23/10/2022",
		},
	}
	errorEsperado := errors.New("Soy un error de Write")
	mockStorage := MockDB{
		dataMock:   database,
		errOnRead:  nil,
		errOnWrite: errors.New("Soy un error de Write"),
		readCheck:  false,
	}
	repo := NewRepository(&mockStorage)
	service := NewService(repo)

	tUpdated := Transaction{
		ID:       1,
		Codigo:   "AFTER UPDATE",
		Moneda:   "USD AFTER",
		Monto:    20.00,
		Emisor:   "Juan",
		Receptor: "Pedro",
		Fecha:    "24/10/2022",
	}

	//act

	result, err := service.Update(tUpdated.ID, tUpdated.Codigo, tUpdated.Moneda,
		tUpdated.Monto, tUpdated.Emisor, tUpdated.Receptor, tUpdated.Fecha)

	//assert

	assert.True(t, mockStorage.readCheck) // SI utiliza el metodo Read dentro de Update entonces deberia levantarse la bandera
	assert.Equal(t, errorEsperado, err)
	assert.Empty(t, result)

}
