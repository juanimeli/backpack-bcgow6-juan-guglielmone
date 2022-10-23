package transactions

import (
	"fmt"

	"github.com/juanimeli/backpack-bcgow6-juan-guglielmone/goTesting/goWebTesting/pkg/store"
)

type Transaction struct {
	ID       int     `json:"ID"`
	Codigo   string  `json:"cod" binding:"required"`
	Moneda   string  `json:"currency" binding:"required"`
	Monto    float64 `json:"amount" binding:"required"`
	Emisor   string  `json:"sender" binding:"required"`
	Receptor string  `json:"receiver" binding:"required"`
	Fecha    string  `json:"date" binding:"required"`
}

type Repository interface {
	GetAll() ([]Transaction, error)
	Store(ID int, cod, currency string, amount float64, sender, receiver, date string) (Transaction, error)
	LastID() (int, error)
	Update(ID int, cod, currency string, amount float64, sender, receiver, date string) (Transaction, error)
	Delete(ID int) error
	UpdateCodnAmount(ID int, cod string, amount float64) (Transaction, error)
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Store(ID int, cod, currency string, amount float64, sender, receiver, date string) (Transaction, error) {

	var ts []Transaction
	if err := r.db.Read(&ts); err != nil {
		return Transaction{}, err
	}
	t := Transaction{ID, cod, currency, amount, sender, receiver, date}
	ts = append(ts, t)

	if err := r.db.Write(ts); err != nil {
		return Transaction{}, err
	}
	return t, nil
}

func (r *repository) GetAll() ([]Transaction, error) {
	var ts []Transaction
	if err := r.db.Read(&ts); err != nil {
		return ts, err
	}
	return ts, nil
}

func (r *repository) Update(ID int, cod, currency string, amount float64, sender, receiver, date string) (Transaction, error) {

	t := Transaction{Codigo: cod, Moneda: currency, Monto: amount, Emisor: sender, Receptor: receiver, Fecha: date}
	updated := false
	var ts []Transaction
	if err := r.db.Read(&ts); err != nil {
		return Transaction{}, err
	}

	for i := range ts {
		if ts[i].ID == ID {
			t.ID = ID
			ts[i] = t
			updated = true
		}
	}
	if !updated {
		return Transaction{}, fmt.Errorf("transaction with id: %d not found", ID)
	}

	if err := r.db.Write(ts); err != nil {
		return Transaction{}, err
	}
	return t, nil
}

func (r *repository) Delete(ID int) error {
	var ts []Transaction
	if err := r.db.Read(&ts); err != nil {
		return err
	}

	deleted := false
	var index int
	for i := range ts {
		if ts[i].ID == ID {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("transaction with id %d not found", ID)
	}
	ts = append(ts[:index], ts[index+1:]...)

	if err := r.db.Write(ts); err != nil {
		return err
	}
	return nil
}

func (r *repository) UpdateCodnAmount(ID int, cod string, amount float64) (Transaction, error) {
	var t Transaction
	var ts []Transaction
	if err := r.db.Read(&ts); err != nil {
		return Transaction{}, fmt.Errorf("somthing went wrong reading the file")
	}
	updated := false
	for i := range ts {
		if ts[i].ID == ID {
			ts[i].Codigo = cod
			ts[i].Monto = amount
			t = ts[i]
			updated = true
		}
	}
	if !updated {
		return Transaction{}, fmt.Errorf("Transaction with id %d not found", ID)
	}

	if err := r.db.Write(ts); err != nil {
		return Transaction{}, fmt.Errorf("error: writing db file")
	}
	return t, nil
}

func (r *repository) LastID() (int, error) {
	var ts []Transaction
	if err := r.db.Read(&ts); err != nil {
		return 0, err
	}
	if len(ts) == 0 {
		return 0, nil
	}
	return ts[len(ts)-1].ID, nil
}
