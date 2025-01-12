package main

import (
	"fmt"
	"io"
	"os"

	"github.com/martishin/load-balancer/internal/adapters"
	"github.com/martishin/load-balancer/internal/domain"
)

func main() {
	orderReader := adapters.NewOrderReader(os.Stdin)

	orderBook := domain.NewInMemoryBook()

	for {
		order, err := orderReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "error reading order: %v\n", err)
			continue
		}

		matches := orderBook.AddOrder(order)

		for _, match := range matches {
			adapters.PrintMatch(os.Stdout, match)
		}
	}
	adapters.PrintBook(os.Stdout, orderBook)
}
