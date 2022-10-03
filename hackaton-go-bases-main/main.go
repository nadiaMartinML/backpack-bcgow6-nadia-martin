package main

import "github.com/nadiaMartinML/backpack-bcgow6-nadia-martin/hackaton-go-bases-main/internal/service"

func main() {
	var tickets []service.Ticket
	// Funcion para obtener tickets del archivo csv
	service.NewBookings(tickets)
}
