package service

import (
	"errors"
	"fmt"
)

type Bookings interface {
	// Create create a new Ticket
	Create(t Ticket) (Ticket, error)
	// Read read a Ticket by id
	Read(id int) (Ticket, error)
	// Update update values of a Ticket
	Update(id int, t Ticket) (Ticket, error)
	// Delete delete a Ticket by id
	Delete(id int) (int, error)
}

type bookings struct {
	Tickets []Ticket
}

type Ticket struct {
	Id                              int
	Names, Email, Destination, Date string
	Price                           int
}

func IDGenerator(currentID int) (newIDint int, newID string) {
	newIDint = currentID + 1
	newID = fmt.Sprintf("%d", newIDint)
	return
}

func (b *bookings) indexOf(id int) (int, error) {
	for i, ticket := range b.Tickets {
		if ticket.Id == id {
			return i, nil
		}
	}
	return 0, errors.New("id not found")
}

// NewBookings creates a new bookings service
func NewBookings(Tickets []Ticket) Bookings {

	return &bookings{Tickets: Tickets}
}

func (b *bookings) Create(t Ticket) (Ticket, error) {
	//chequear si existe el id que se esta ingresando, hecho
	//sino asignarle uno con un generador de id o lanzar error
	//Crear nuevo ticket con id que se paso o id auto generado

	for _, ticket := range b.Tickets {
		if ticket.Id == t.Id {
			return ticket, errors.New("the ID alrready exists")
		}
	}

	b.Tickets = append(b.Tickets, t)

	return t, nil
}

func (b *bookings) Read(id int) (Ticket, error) {
	for _, ticket := range b.Tickets {
		if ticket.Id == id {
			return ticket, nil
		}
	}
	return Ticket{}, errors.New("ticket not found, try another id")
}

func (b *bookings) Update(id int, t Ticket) (Ticket, error) {

	ticket, err := b.Read(id)
	if err != nil {
		return ticket, err
	}

	i, err := b.indexOf(id)

	if err != nil {
		return ticket, err
	}

	b.Tickets[i] = t

	return Ticket{}, nil
}

func (b *bookings) Delete(id int) (int, error) {

	i, err := b.indexOf(id)

	if err != nil {
		return 0, err
	}
	b.Tickets = append(b.Tickets[:i], b.Tickets[i+1:]...)
	fmt.Printf("El ticket con id %d ha sido eliminado correctamente.\n", id)
	return 1, nil
}
