package main

import (
	"github.com/honey-yogurt/yolog"
	"log"
	"os"
)

func main() {
	fd, err := os.OpenFile("test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln("create file test.log failed")
	}
	defer fd.Close()

	logger := yolog.New(yolog.WithLevel(yolog.DebugLevel), yolog.WithOutput(fd), yolog.WithFormatter(&yolog.JsonFormatter{IgnoreBasicFields: false}))
	logger.Infof("this is logger %v", logger)
}
