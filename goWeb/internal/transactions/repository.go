package transactions

/*
Repositorio, debe tener el acceso a la variable guardada en memoria.
Se debe crear el archivo repository.go
Se debe crear la estructura de la entidad
Se deben crear las variables globales donde guardar las entidades
Se debe generar la interface Repository con todos sus métodos
Se debe generar la estructura repository
Se debe generar una función que devuelva el Repositorio
Se deben implementar todos los métodos correspondientes a las operaciones a realizar (GetAll, Store, etc..)
*/

type Transaction struct {
	ID       int     `json:"ID"`
	Codigo   string  `json:"cod" binding:"required"`
	Moneda   string  `json:"currency" binding:"required"`
	Monto    float64 `json:"amount" binding:"required"`
	Emisor   string  `json:"sender" binding:"required"`
	Receptor string  `json:"receiver" binding:"required"`
	Fecha    string  `json:"date" binding:"required"`
}

var ts []Transaction
var lastID int

type Repository interface {
	GetAll() ([]Transaction, error)
	AddTransaction(ID int, cod, currency string, amount float64, sender, receiver, date string) (Transaction, error)
	LastID() (int, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) AddTransaction(ID int, cod, currency string, amount float64, sender, receiver, date string) (Transaction, error) {

	t := Transaction{ID, cod, currency, amount, sender, receiver, date}
	ts = append(ts, t)
	lastID = t.ID // actualiza el lastID global cuando se va a agregar una transaccion
	return t, nil
}

func (r *repository) GetAll() ([]Transaction, error) {
	return ts, nil
}

func (r *repository) LastID() (int, error) {
	return lastID, nil
}
