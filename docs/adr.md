## Logging 

Logging should take in periodic health checks for container resources, availability, throughput, latency, and other performance metrics. 

Logging should have seperate structures to log different metrics. This enables a way to track, 
aggregate, and provide reports periodically. 

Logs should be tagged with a timestamp & have consistent formatting to enable for
a modular mechanism for logs. 

More on the topic of performance, we should be tracking latency with OpenTelemetry as it provides a way to split total latency up by each process helping us identify bottlenecks. 

## API Entrypoint

The 

