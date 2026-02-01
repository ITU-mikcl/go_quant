# Go Quant

A CLI tool written in Go that calculates the 10-year **Total Return** and **Annualized Sharpe Ratio** for a given stock symbol using the Alpaca Markets API.

## Setup

### Configure Environment Variables:

Create a .env file in the root directory, and add your Alpaca credentials to the file:

```
APCA_API_KEY_ID=your_key_here
APCA_API_SECRET_KEY=your_secret_here
```

## Usage

Run the program by passing a stock ticker as the first argument:

```bash
go run main.go GOOGL
```