package main

import (
	"context"
	"encoding/json"
	"log"

	alpacaconsumer "github.com/Quantlytic/AlpacaConsumer/internal/AlpacaConsumer"
	"github.com/Quantlytic/AlpacaConsumer/internal/config"
	"github.com/Quantlytic/AlpacaConsumer/pkg/kafkaproducer"
	"github.com/alpacahq/alpaca-trade-api-go/v3/marketdata/stream"
)

func main() {
	// Load environment variables
	appConfig := config.Load()

	// Initialize consumer
	consumer, err := alpacaconsumer.NewAlpacaConsumer(alpacaconsumer.AlpacaConsumerConfig{
		Stream:  "test",
		ApiKey:  appConfig.AlpacaAPIKey,
		Secret:  appConfig.AlpacaSecretKey,
		BaseURL: appConfig.AlpacaBaseURL,
	})

	if err != nil {
		log.Fatalf("Error creating Alpaca consumer: %v", err)
	}

	ctx := context.Background()

	// Establish websocket connection with API
	if err := consumer.Connect(ctx); err != nil {
		log.Fatalf("Error connecting to Alpaca: %v", err)
	}

	p, err := kafkaproducer.NewKafkaProducer(kafkaproducer.DefaultConfig)
	if err != nil {
		log.Fatalf("Error creating Kafka producer: %v", err)
	}

	// Helper function to handle JSON marshaling and Kafka producing
	produceToKafka := func(topic, symbol string, data interface{}) {
		dataJSON, err := json.Marshal(data)
		if err != nil {
			log.Printf("Error marshaling data to JSON: %v", err)
			return
		}
		p.Produce(topic, []byte(symbol), dataJSON)
	}

	// Array of symbols
	// symbols := []string{"MSFT", "AAPL", "GOOG", "AMZN", "META"}
	symbols := []string{"FAKEPACA"}

	// Subscribe to quotes
	if err := consumer.SubscribeToQuotes(ctx, symbols, func(quote stream.Quote) {
		log.Printf("Received quote: %+v", quote)
		produceToKafka("stock-data-raw", quote.Symbol, quote)
	}); err != nil {
		log.Fatalf("Error subscribing to symbols: %v", err)
	}

	// Subscribe to trades
	if err := consumer.SubscribeToTrades(ctx, symbols, func(trade stream.Trade) {
		log.Printf("Received trade: %+v", trade)
		produceToKafka("stock-data-raw", trade.Symbol, trade)
	}); err != nil {
		log.Fatalf("Error subscribing to trades: %v", err)
	}

	<-consumer.Client.Terminated()
}
