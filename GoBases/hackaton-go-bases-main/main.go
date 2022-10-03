package main

import (
	"fmt"
	"hackaton/internal/file"
	"hackaton/internal/service"
)

func main() {

	var tickets []service.Ticket
	// Funcion para obtener tickets del archivo csv
	service.NewBookings(tickets)

	var archivo file.File = file.File{"./tickets.csv"}

	ticketsOnFile, _ := archivo.Read()
	fmt.Println(ticketsOnFile[489])

	bookings := service.NewBookings(ticketsOnFile)

	ticket, _ := bookings.Read(15)

	fmt.Println(ticket)

	ticketUpdate := service.Ticket{
		Id:          15,
		Names:       "Juan Gugli",
		Email:       "juan@meli",
		Destination: "uruguay",
		Date:        "19 setiembre",
		Price:       696,
	}
	bookings.Update(15, ticketUpdate)
	ticketUpdated, _ := bookings.Read(15)
	fmt.Println(ticketUpdated)

	bookings.Delete(378)
	/*ticketAfter, errDe := bookings.Read(283)

	if errDe != nil {
		fmt.Println(errDe)
		os.Exit(1)
	}
	fmt.Println(ticketAfter)*/

	newTicket := service.Ticket{
		Id:          2000,
		Names:       "Juan Guasdasgli",
		Email:       "juan@asdasdmeli",
		Destination: "uruasdasdguay",
		Date:        "19 seasdasdtiembre",
		Price:       6234236,
	}

	_, err3 := bookings.Create(newTicket)

	if err3 != nil {
		fmt.Println(err3)
	}

}
