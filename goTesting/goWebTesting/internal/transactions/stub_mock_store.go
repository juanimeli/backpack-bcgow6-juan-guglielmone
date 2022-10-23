package transactions

// En este caso creo un StubDB que va a imitar mi fileStore de store/file.go. Esta estructura StubDB NO va a
// ser igual a la struct fileStore. Esto es porque
// lo que queremos simular debe cumplir con el tipo de interface Store que es la requerida para
// crear el repository que es lo que queremos testear, en particular el metodo GetAll(). Para esto creamos
// la estructura StubDB con los metodos que cumplan el contrato de interface Store. Dentro del metodo
// Read que imita el metodo Read de fileStore, primero declaramos que la interface que estamos pasando
// es del tipo Slice de Transaction y luego como que le asigna a la interface que se le pasa a Read() en
// GetAll() la dataStub del StubDB que es lo que definimos como la database ficticia.

type StubDB struct {
	dataStub []Transaction
}

func (stb StubDB) Read(data interface{}) error {

	castedData := data.(*[]Transaction)
	*castedData = stb.dataStub

	return nil
}

func (std StubDB) Write(data interface{}) error {
	return nil
}

// Para el ejercico 2 de la calse 2 TM realizamos un Mock ya que ademas de probar que nos devuelve la
// transaction updated, queremos chequear que se llamo la funcion Read(), esto lo haremos colocando
// una bandera en el metodo Read() del Mock.

type MockDB struct {
	readCheck  bool
	DataMock   []Transaction
	errOnRead  error
	errOnWrite error
}

func (m *MockDB) Read(data interface{}) error {
	if m.errOnRead != nil {
		return m.errOnRead
	}

	castedData := data.(*[]Transaction)
	*castedData = m.DataMock
	m.readCheck = true // bandera que indica que el metodo Read() fue llamado en el proceso de Update()
	return nil
}

// el metodo write simplemente guarda de nuevo la interface con los datos actualizados en la
// logica de Update.

func (m *MockDB) Write(data interface{}) error {
	if m.errOnWrite != nil {
		return m.errOnWrite
	}
	castedData := data.([]Transaction)
	m.DataMock = castedData
	return nil
}
