package main

import (
	"fmt"
	"github.com/t-k/fluent-logger-golang/fluent"
	"os"
	"strconv"
)

func main() {
	tagPrefix := os.Getenv("FLUENT_TAG")
	if tagPrefix == "" {
		tagPrefix = "asterisk"
	}
	host := os.Getenv("FLUENT_HOST")
	if host == "" {
		host = "localhost"
	}
	portStr := os.Getenv("FLUENT_PORT")
	if portStr == "" {
		portStr = "24224"
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	logger, err := fluent.New(fluent.Config{FluentPort: port, FluentHost: host})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer logger.Close()

	tag := os.Args[1]
	message := os.Args[2:]
	if tag != "" && message != nil {
		data := map[string]interface{}{"message": message}
		logger.Post(tagPrefix+"."+tag, data)
	}
}
