package transactions

type Service interface {
	GetAll() ([]Transaction, error)
	Store(cod, currency string, amount float64, sender, receiver, date string) (Transaction, error)
	Update(ID int, cod, currency string, amount float64, sender, receiver, date string) (Transaction, error)
	Delete(Id int) error
	UpdateCodnAmount(ID int, cod string, amount float64) (Transaction, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]Transaction, error) {
	t, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (s *service) Store(cod, currency string, amount float64, sender, receiver, date string) (Transaction, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return Transaction{}, err
	}
	lastID++

	transaction, err := s.repository.Store(lastID, cod, currency, amount, sender, receiver, date)
	if err != nil {
		return Transaction{}, err
	}
	return transaction, nil
}

func (s *service) Update(ID int, cod, currency string, amount float64, sender, receiver, date string) (Transaction, error) {
	return s.repository.Update(ID, cod, currency, amount, sender, receiver, date)
}

func (s *service) Delete(ID int) error {
	return s.repository.Delete(ID)
}

func (s *service) UpdateCodnAmount(ID int, cod string, amount float64) (Transaction, error) {
	return s.repository.UpdateCodnAmount(ID, cod, amount)
}
