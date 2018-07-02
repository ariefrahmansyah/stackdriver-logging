# stackdriver-logging

Stackdriver Logging is wrapper of [Google's Stackdriver Loogging](https://cloud.google.com/logging/). It's heavily inspired by [Prometheus' logging](https://github.com/prometheus/common/blob/master/log).

## Usage
```go
ctx := context.Background()
projectID := "my-awesome-project"

// Initializes logging client
loggingClient, err := logging.NewClient(ctx, projectID)
if err != nil {
    panic(err)
}

// Initializes stackdriver-logging
log.InitLogger(loggingClient, "spinnaker-demo-logger")

// Start logging. Your logging will show up at your Stackdriver Logs
log.Info("This is info")
log.Notice("This is notice")
log.Warning("This is warning")
log.Error("This is error")
log.Critical("This is critical")
log.Alert("This is alert")
log.Emergency("This is emergency")
```
