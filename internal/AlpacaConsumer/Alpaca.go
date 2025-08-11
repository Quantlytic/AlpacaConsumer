package alpacaconsumer

import (
	"fmt"
	"context"
	"github.com/alpacahq/alpaca-trade-api-go/v3/marketdata"
	"github.com/alpacahq/alpaca-trade-api-go/v3/marketdata/stream"
)

type OnMarketData func()

type AlpacaConsumer struct {
	Client *stream.StocksClient
}

func NewAlpacaConsumer(apiKey, secretKey, baseUrl string) (*AlpacaConsumer, error) {
	c := stream.NewStocksClient(
		marketdata.IEX,
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
func (ac *AlpacaConsumer) connect(ctx context.Context) error {
	if err := ac.Client.Connect(ctx); err != nil {
		return fmt.Errorf("failed to connect to Alpaca market data stream: %w", err)
	}
	return nil
}

// SubscribeToSymbols subscribes to the given symbols for market data updates.
func (ac *AlpacaConsumer) SubscribeToSymbols(ctx context.Context, symbols []string) error {
	if err := ac.Client.SubscribeToTrades(ctx, symbols...); err != nil {
		return fmt.Errorf("failed to subscribe to symbols: %w", err)
	}
	return nil
}