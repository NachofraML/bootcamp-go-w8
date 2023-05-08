package tickets

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTicketsPackage(t *testing.T) {
	//Setting path to csv file from test module
	CsvFilePath = "/Users/imosca/Documents/Bootcamp GO/01-go-bases/clase-05/desafio-go-bases/tickets.csv"

	t.Run("TestGetTotalTickets", func(t *testing.T) {
		var expectedTotalTickets = 45
		total, err := GetTotalTickets("Brazil")
		assert.Nil(t, err)
		assert.Equal(t, expectedTotalTickets, total)
	})

	t.Run("TestGetCountByPeriod", func(t *testing.T) {
		var period = "dawn"
		var expectedTotalTickets = 304
		totalCountByPeriod, err := GetCountByPeriod(period)
		assert.Nil(t, err)
		assert.Equal(t, expectedTotalTickets, totalCountByPeriod)
	})

	t.Run("TestAverageDestination", func(t *testing.T) {
		var destination = "Brazil"
		var total = 100.0
		var expectedAverage = 45
		actualAverage, err := AverageDestination(destination, total)
		assert.Nil(t, err)
		assert.Equal(t, expectedAverage, actualAverage)
	})
}