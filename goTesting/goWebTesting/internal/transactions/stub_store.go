package transactions

type StubDB struct {
	dataStub []Transaction
}

func (stb StubDB) Read(data interface{}) error {

	casteData := data.(*[]Transaction)
	*casteData = stb.dataStub

	return nil
}

func (std StubDB) Write(data interface{}) error {

	castedData := data.(*Transaction)
	std.dataStub = append(std.dataStub, *castedData)
	return nil
}
