package main

import (
	"context"
	"log"

	alpacaconsumer "github.com/Quantlytic/AlpacaConsumer/internal/AlpacaConsumer"
	"github.com/Quantlytic/AlpacaConsumer/internal/config"
	"github.com/alpacahq/alpaca-trade-api-go/v3/marketdata/stream"
)

func main() {
	// Load environment variables
	appConfig := config.Load()

	// Initialize consumer
	consumer, err := alpacaconsumer.NewAlpacaConsumer(
		appConfig.AlpacaAPIKey,
		appConfig.AlpacaSecretKey,
		appConfig.AlpacaBaseURL,
	)

	if err != nil {
		log.Fatalf("Error creating Alpaca consumer: %v", err)
	}

	ctx := context.Background()

	// Establish websocket connection with API
	if err := consumer.Connect(ctx); err != nil {
		log.Fatalf("Error connecting to Alpaca: %v", err)
	}

	// Array of symbols
	symbols := []string{"FAKEPACA"}

	// Subscribe to quotes
	if err := consumer.SubscribeToQuotes(ctx, symbols, func(quote stream.Quote) {
		log.Printf("Received quote: %+v", quote)
	}); err != nil {
		log.Fatalf("Error subscribing to symbols: %v", err)
	}

	// Subscribe to trades
	if err := consumer.SubscribeToTrades(ctx, symbols, func(trade stream.Trade) {
		log.Printf("Received trade: %+v", trade)
	}); err != nil {
		log.Fatalf("Error subscribing to trades: %v", err)
	}

	<-consumer.Client.Terminated()
}
