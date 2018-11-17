package main

import (
	"time"

	"github.com/go-apibox/logging"
)

func main() {
	go func() {
		test("test1")
	}()

	test("test2")

	time.Sleep(time.Second)
}

func test(name string) {
	logger := logging.NewLogger(name)
	logger.Debug("debug")
	logger.Info("info")
	logger.Notice("notice")
	logger.Warning("warning")
	logger.Error("error")
	logger.Critical("critical")
}
