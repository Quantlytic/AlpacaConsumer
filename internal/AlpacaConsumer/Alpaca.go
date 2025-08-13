package alpacaconsumer

import (
	"context"
	"fmt"

	"github.com/alpacahq/alpaca-trade-api-go/v3/marketdata/stream"
)

type OnQuote func(stream.Quote)
type OnTrade func(stream.Trade)

type AlpacaConsumer struct {
	Client *stream.StocksClient
}

func NewAlpacaConsumer(apiKey, secretKey, baseUrl string) (*AlpacaConsumer, error) {
	// Debug print function inputs
	fmt.Printf("NewAlpacaConsumer called with apiKey: %s, secretKey: %s, baseUrl: %s\n", apiKey, secretKey, baseUrl)

	c := stream.NewStocksClient(
		"test",
		stream.WithCredentials(apiKey, secretKey),
		stream.WithBaseURL(baseUrl),
	)

	// Check if the client is properly initialized
	if c == nil {
		return nil, fmt.Errorf("failed to create Alpaca market data client")
	}
	return &AlpacaConsumer{Client: c}, nil
}

// Connect establishes a connection to the Alpaca market data stream.
func (ac *AlpacaConsumer) Connect(ctx context.Context) error {
	if err := ac.Client.Connect(ctx); err != nil {
		return fmt.Errorf("failed to connect to Alpaca market data stream: %w", err)
	}

	return nil
}

// SubscribeToSymbols subscribes to the given symbols for market data updates.
func (ac *AlpacaConsumer) SubscribeToQuotes(ctx context.Context, symbols []string, handler OnQuote) error {
	if err := ac.Client.SubscribeToQuotes(handler, symbols...); err != nil {
		return fmt.Errorf("failed to subscribe to symbols: %w", err)
	}
	return nil
}

func (ac *AlpacaConsumer) SubscribeToTrades(ctx context.Context, symbols []string, handler OnTrade) error {
	if err := ac.Client.SubscribeToTrades(handler, symbols...); err != nil {
		return fmt.Errorf("failed to subscribe to symbols: %w", err)
	}
	return nil
}
