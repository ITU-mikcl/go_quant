package main

import (
	"fmt"
	"os"
	"time"

	"github.com/alpacahq/alpaca-trade-api-go/v3/marketdata"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	client := marketdata.NewClient(marketdata.ClientOpts{
		APIKey:    os.Getenv("APCA_API_KEY_ID"),
		APISecret: os.Getenv("APCA_API_SECRET_KEY"),
		BaseURL:   "https://data.alpaca.markets",
	})

	latestTradingDay := getLatestTradingDay()

	request := marketdata.GetBarsRequest{
		TimeFrame: marketdata.OneDay,
		Start:     latestTradingDay.AddDate(-10, 0, 0),
		End:       latestTradingDay,
	}

	bars, err := client.GetBars("GOOGL", request)
	if err != nil {
		panic(err)
	}
	for _, bar := range bars {
		fmt.Printf("%+v\n", bar)
	}
}

// getLatestTradingDay returns today if it's a weekday, or the last Friday if today is a weekend
func getLatestTradingDay() time.Time {
	now := time.Now()
	weekday := now.Weekday()

	// If today is Saturday or Sunday, go back to Friday
	switch weekday {
	case time.Saturday:
		return now.AddDate(0, 0, -1)
	case time.Sunday:
		return now.AddDate(0, 0, -2)
	}

	// Otherwise, today is a trading day (Mon-Fri)
	return now
}
