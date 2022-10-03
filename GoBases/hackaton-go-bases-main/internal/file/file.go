package file

import (
	"fmt"
	"hackaton/internal/service"
	"os"
	"strconv"
	"strings"
)

type File struct {
	Path string
}

func (f *File) Read() ([]service.Ticket, error) {

	data, err := os.ReadFile(f.Path)
	if err != nil {
		return nil, err
	}

	dataS := string(data)

	lines := strings.Split(dataS, "\n")
	fmt.Println(len(lines))
	fmt.Println(lines[1])

	var ticketsOnFile []service.Ticket

	for _, line := range lines {
		if len(line) > 0 {

			dataLine := strings.Split(line, ",")
			idInt, _ := strconv.Atoi(dataLine[0])
			priceInt, _ := strconv.Atoi(dataLine[5])

			ticket := service.Ticket{idInt, dataLine[1], dataLine[2], dataLine[3], dataLine[4], priceInt}
			ticketsOnFile = append(ticketsOnFile, ticket)

		}
	}
	return ticketsOnFile, err
}

func (f *File) Write(ticket *service.Ticket) (err error) {

	newTicket := fmt.Sprintf("\n%d,%s,%s,%s,%s,%d",
		ticket.Id,
		ticket.Names,
		ticket.Email,
		ticket.Destination,
		ticket.Price,
	)

	file, err := os.OpenFile(f.Path, os.O_WRONLY|os.O_APPEND, 0666)

	defer file.Close()

	if err != nil {
		return
	}

	_, err2 := file.Write([]byte(newTicket))

	if err2 != nil {
		return err2
	}

	file.Close()
	return nil
}

/* func IDExist(file, id string) (res bool, err error) {
	f, err := os.Open(file)
	if err != nil {
		return
	}

	defer f.Close()

	//dataS := string(data)

	r := csv.NewReader(f)

	for {
		record, err := r.Read()
		//fmt.Println(record)

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if record[0] == id {
			res = true
			break
		}

	}
	return

}
*/
