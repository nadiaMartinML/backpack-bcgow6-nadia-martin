package file

import "github.com/nadiaMartinML/backpack-bcgow6-nadia-martin/hackaton-go-bases-main/internal/service"

type File struct {
	path string
}

func (f *File) Read() ([]service.Ticket, error) {
	return nil, nil
}

func (f *File) Write(service.Ticket) error {
	return nil
}
