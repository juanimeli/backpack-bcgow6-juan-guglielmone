package transactions

type StubDB struct {
	dataStub []Transaction
}

func (stb StubDB) Read(data interface{}) error {

	/*
		transactions := []Transaction{
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
	*/
	casteData := data.(*[]Transaction)
	*casteData = stb.dataStub

	return nil
}

/*
	func (d StubDB) buscarPorCodigo(codigo string) string {
		return ""
	}
*/
func (std StubDB) Write(data interface{}) error {

	castedData := data.(*Transaction)
	std.dataStub = append(std.dataStub, *castedData)
	return nil
}
