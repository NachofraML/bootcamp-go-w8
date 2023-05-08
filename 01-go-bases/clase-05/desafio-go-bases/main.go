package main

import (
	"fmt"
	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

func main() {
	totalCountryTickets, err := tickets.GetTotalTickets("Brazil")
	totalCountByPeriod, err := tickets.GetCountByPeriod("dawn")
	average, err := tickets.AverageDestination("Brazil", 2000)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Total tickets to Brazil: %d\n", totalCountryTickets)
	fmt.Printf("Total tickets by dawn: %d\n", totalCountByPeriod)
	fmt.Printf("Percentage of brazilians tickets in base of total tickets: %d\n", average)
}
