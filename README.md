# Alpaca Consumer
The Alpaca Consumer establishes a connection with the Alpaca trading API to receive live market data. The free tier of this API is limited to 30 symbols and to the IEX exchange only. After this service is live, it will be possible to test the full data collection pipeline.

## Service Requirements
1. Establish websocket connection to Alpaca's API
2. Stream live data from Alpaca API
3. Write data to stock-data-raw kafka topic

## Goals
* When writing to kafka topic, partition data by symbol to allow parallelization and batching on the kafka consumer side
