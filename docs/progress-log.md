# Project Progress / Notes

**August 12th -** Wrote a first draft of the project which subscribes to the Alpaca API live testing websocket and prints the data to the command line. After confirming that my credentials worked and I could stream data, I wrote a kafka producer that could write the live data to the kafka topic.

Changes that need to be made before collecting live data:
* Setup build / deployment to kubernetes
* Configure batching / file storage in db write service
* Conform messages to write service's expected json format
* Architecture decisions on trade / quote data
  * Should they use the same topic or different topics? They have different parameters

