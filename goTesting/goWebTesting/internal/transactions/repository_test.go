package transactions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Para testear el GetAll con el Stub en Arrange defino la base de datos ficticia que sera guardada
// en la StudDB (doble de base de datos) que fue creada para satisfacer la dependencia de NewRepository
// y asi probar el metrodo GetAll() que en este test debera leer esta base de datos ficticia. Lo que esta
// sucediendo es que una vez creado el repository con la StubDB que imita la fileStore, se llama al metodo
// GetAll() del repositorio, el cual esta siendo testeado y en su logica va a llamar a el metodo Read(),
// que como el repo esta instanciado con el StubDB usara el metodo de esta estructura ultimo mencionada.
// Este metodo read se encarga de meter la data definida en la estructura (base de datos ficticia
// dentro de la interface que devolvera el metodo GetAll() del repository. Luego solamente comparamos
// que la database ficticia coincida con la leida por el metodo GetAll() que leera de una base de datos
// incluida en el doble Stub el cual imita el fileStore.
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

	//result = append(result, Transaction{})
	// Agrega una Transaction vacia al resultado de GetAll para que no coincida
	// con la lista de Transactions de la db y falle el Test

	//assert
	assert.Nil(t, err)
	assert.Equal(t, database, result)
}

func TestUpdateCodnAmount(t *testing.T) {

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

	MockStorage := MockDB{
		dataMock:  database,
		readCheck: false,
	}
	repo := NewRepository(&MockStorage)

	expected := Transaction{
		ID:       1,
		Codigo:   "AFTER UPDATE",
		Moneda:   "USD",
		Monto:    20.00,
		Emisor:   "Juan",
		Receptor: "Pedro",
		Fecha:    "23/10/2022",
	}

	//act

	result, err := repo.UpdateCodnAmount(1, "AFTER UPDATE", 20.00)
	assert.Nil(t, err) //Cuando llamaos a la funcion UpdateCodnAmount vamos a utilizar el recurso de Update
	// del repo con la base de datos ficticia guardada en nuestor Mock. Basicamente hace lo mismo que
	// en el stub pero se le suma que cuando el metodo Update() llama a la funcion Read() se levanta la
	// bandera de true y se guarda en la estructura del doble Mock.

	assert.True(t, MockStorage.readCheck) // comprobamos que la bandera se haya levantado en el metodo Read()
	// de nuestro Mock
	assert.Equal(t, expected, result) // se puede hacer dos Equal y solo chequear los campos result.Codigo & result.Monto.
	// En mi caso iguale las tracciones en su totalidad

}

func TestStore(t *testing.T) {

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

	//act

	result, err := repo.Store(3, "NEW", "UYU", 20.00, "Juan", "Martin", "22/10")

	//assert
	assert.Nil(t, err)
	assert.Equal(t, esperado, result)
}

func TestUpdate(t *testing.T) {
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

	MockStorage := MockDB{
		dataMock:   database,
		readCheck:  false,
		errOnRead:  nil,
		errOnWrite: nil,
	}
	repo := NewRepository(&MockStorage)

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

	result, err := repo.Update(1, "AFTER UPDATE", "USD AFTER", 20.00, "Juan", "Pedro", "24/10/2022")

	// assert
	assert.Nil(t, err)

	assert.True(t, MockStorage.readCheck)
	assert.Equal(t, expected, result)
}

func TestDelete(t *testing.T) {
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

	MockStorage := MockDB{
		dataMock:   database,
		readCheck:  false,
		errOnRead:  nil,
		errOnWrite: nil,
	}
	repo := NewRepository(&MockStorage)

	//act

	err := repo.Delete(1)

	// assert
	assert.Nil(t, err)
	assert.True(t, MockStorage.readCheck)
	assert.Equal(t, 1, len(MockStorage.dataMock))

}
