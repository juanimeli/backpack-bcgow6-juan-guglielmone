package file

import (
	"hackaton/internal/service"
	"os"
	"fmt"
)

type File struct {
	path string
}

func (f *File) Read() ([]service.Ticket, error) {

	data, err := os.ReadFile()


	return nil, nil
}

func (f *File) Write(service.Ticket) error {
	return nil
}
