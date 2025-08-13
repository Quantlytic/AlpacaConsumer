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

type AlpacaConsumerConfig struct {
	Stream  string
	ApiKey  string
	Secret  string
	BaseURL string
}

func NewAlpacaConsumer(config AlpacaConsumerConfig) (*AlpacaConsumer, error) {
	c := stream.NewStocksClient(
		config.Stream,
		stream.WithCredentials(config.ApiKey, config.Secret),
		stream.WithBaseURL(config.BaseURL),
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
