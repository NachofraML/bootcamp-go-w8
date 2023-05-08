package tickets

import (
	"github.com/gocarina/gocsv"
	"log"
	"os"
	"strings"
	"time"
)

type Ticket struct {
	ID            int64  `csv:"id"`
	FullName      string `csv:"full_name"`
	Email         string `csv:"email"`
	Country       string `csv:"country"`
	BoughtAt      string `csv:"buyed_at"`
	IdkWhatIsThis int    `csv:"amount"`
}

var (
	CsvFilePath = "./tickets.csv"
)

func readFile(path string) (file *os.File, err error) {
	file, err = os.Open(path)
	if err != nil {
		log.Fatal("Unable to read input file", err)
	}
	return
}

func csvFileToTickets(file *os.File) (tickets []Ticket, err error) {
	tickets = []Ticket{}
	if err = gocsv.UnmarshalFile(file, &tickets); err != nil {
		log.Fatal("Unable to parse CSV files to Tickets", err)
	}
	return
}

// ejemplo 1
func GetTotalTickets(destination string) (totalTicketsCount int, err error) {
	file, err := readFile(CsvFilePath)
	defer file.Close()
	tickets, err := csvFileToTickets(file)
	if err != nil {
		return
	}
	totalTicketsCount = 0
	for _, ticket := range tickets {
		if ticket.Country == destination {
			totalTicketsCount++
		}
	}
	return
}

var (
	HoursRange = map[string]string{
		"dawn":      "00:00:00;07:00:00",
		"morning":   "07:00:00;13:00:00",
		"afternoon": "13:00:00;20:00:00",
		"night":     "20:00:00;00:00:00",
	}
	HoursLayout = "15:04:05"
)

// ejemplo 2
func GetCountByPeriod(selectedTime string) (totalTicketsCount int, err error) {
	file, err := readFile(CsvFilePath)
	defer file.Close()
	tickets, err := csvFileToTickets(file)
	if err != nil {
		return
	}

	searchedRangeOfHours := HoursRange[selectedTime]
	if searchedRangeOfHours == "" {
		log.Fatal("Unable to find a range of hours for that word", err)
	}

	selectedHours := strings.Split(searchedRangeOfHours, ";")

	startHour, err := time.Parse(HoursLayout, selectedHours[0])
	if err != nil {
		log.Fatal("Unable to parse start hour ", err)
	}
	endHour, err := time.Parse(HoursLayout, selectedHours[1])
	if err != nil {
		log.Fatal("Unable to parse start hour ", err)
	}

	totalTicketsCount = 0
	for _, ticket := range tickets {
		boughtAtParsed, err := time.Parse(HoursLayout, ticket.BoughtAt+":00")
		if err != nil {
			log.Fatal("Unable to parse bought_at ", err)
		}
		if startHour.Before(boughtAtParsed) && endHour.After(boughtAtParsed) {
			totalTicketsCount++
		}
	}
	return
}

// ejemplo 3
func AverageDestination(destination string, total float64) (int, error) {
	totalTicketsOfCountry, err := GetTotalTickets(destination)
	if err != nil {
		return 0, err
	}
	average := int(float64(totalTicketsOfCountry) / total * 100)
	return average, nil
}
