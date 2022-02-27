package main

import (
	"errors"
	"io"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

// https://pkg.go.dev/github.com/go-kit/kit/log/level
// https://blog.advenoh.pe.kr/go/Go%EC%97%90%EC%84%9C%EC%9D%98-%EB%A1%9C%EA%B7%B8%EA%B9%85-Logging-in-Go/
// https://golang.hotexamples.com/examples/github.com.hashicorp.logutils/-/LogLevel/golang-loglevel-function-examples.html

func main() {
	// Set up logger with level filter.
	// logger := log.NewLogfmtLogger(os.Stdout)
	// logger := log.NewLogfmtLogger(ioutil.Discard)

	logFile, _ := os.Create("log.txt")
	// logger := log.NewLogfmtLogger(logFile)
	multiWriter := io.MultiWriter(logFile, os.Stdout)
	// logger := log.NewLogfmtLogger(multiWriter)
	logger := log.NewLogfmtLogger(log.NewSyncWriter(multiWriter)) // logging logic is performed within a mutex

	logger = level.NewFilter(logger, level.AllowInfo())
	logger = log.With(logger, "caller", log.DefaultCaller)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	// Use level helpers to log at different levels.
	level.Error(logger).Log("err", errors.New("bad data"))
	level.Info(logger).Log("event", "data saved")
	level.Debug(logger).Log("next item", 17) // filtered

	logger.Log("foo", "bar") // as normal, no level
}
