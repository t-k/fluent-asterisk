package main

import (
	"fmt"
	"github.com/t-k/fluent-logger-golang/fluent"
	"os"
)

func main() {
	tagPrefix := os.Getenv("FLUENT_TAG")
	if tagPrefix != "" {
		tagPrefix = "asterisk"
	}
	host := os.Getenv("FLUENT_HOST")
	if host != "" {
		host = "localhost"
	}
	port := os.Getenv("FLUENT_PORT")
	if port != "" {
		port = "24224"
	}
	logger, err := fluent.New(fluent.Config{})
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	defer logger.Close()

	tag := os.Args[1]
	message := os.Args[2:]
	if tag != "" && message != nil {
		data := map[string]interface{}{"message": message}
		logger.Post(tag, data)
	}
}
