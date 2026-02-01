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

	symbol := "GOOGL"
	latestTradingDay := getLatestTradingDay()

	request := marketdata.GetBarsRequest{
		TimeFrame:  marketdata.OneDay,
		Start:      latestTradingDay.AddDate(-10, 0, 0),
		End:        latestTradingDay,
		Adjustment: marketdata.All,
	}

	bars, err := client.GetBars(symbol, request)
	if err != nil {
		panic(err)
	}

	startPrice := bars[0].Close
	endPrice := bars[len(bars)-1].Close
	totalReturn := ((endPrice - startPrice) / startPrice) * 100

	fmt.Printf("Ticker:           %s\n", symbol)
	fmt.Printf("Start Date:       %s (Price: $%.2f)\n", bars[0].Timestamp.Format("2006-01-02"), startPrice)
	fmt.Printf("End Date:         %s (Price: $%.2f)\n", bars[len(bars)-1].Timestamp.Format("2006-01-02"), endPrice)
	fmt.Printf("Total Return:     %.2f%%\n", totalReturn)
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
