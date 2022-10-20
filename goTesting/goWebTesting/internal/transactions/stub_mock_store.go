package transactions

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

type MockDB struct {
	readCheck bool
	dataMock  []Transaction
}

func (m *MockDB) Read(data interface{}) error {
	castedData := data.(*[]Transaction)
	*castedData = m.dataMock
	m.readCheck = true
	return nil
}

func (m *MockDB) Write(data interface{}) error {
	castedData := data.([]Transaction)
	m.dataMock = castedData
	return nil
}
