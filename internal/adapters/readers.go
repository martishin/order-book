package adapters

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/martishin/load-balancer/internal/models"
)

const ExpectedPartsCount = 4

type OrderReader struct {
	scanner *bufio.Scanner
}

func (or *OrderReader) Next() (*models.Order, error) {
	if or.scanner.Scan() {
		line := or.scanner.Text()
		parts := strings.Split(line, ",")
		if len(parts) != ExpectedPartsCount {
			return nil, fmt.Errorf("invalid line: %s", line)
		}

		orderID, err := strconv.ParseInt(parts[0], 10, 32)
		if err != nil {
			return nil, fmt.Errorf("invalid order id: %s", line)
		}
		orderType := models.OrderType(parts[1])
		price, err := strconv.ParseInt(parts[2], 10, 32)
		if err != nil {
			return nil, fmt.Errorf("invalid order price: %s", line)
		}
		quantity, err := strconv.ParseInt(parts[3], 10, 32)
		if err != nil {
			return nil, fmt.Errorf("invalid quantity: %s", line)
		}

		return &models.Order{
			ID:       int32(orderID),
			Price:    int32(price),
			Quantity: int32(quantity),
			Type:     orderType,
		}, nil
	}

	if err := or.scanner.Err(); err != nil {
		return nil, err
	}

	return nil, io.EOF
}

func NewOrderReader(r io.Reader) *OrderReader {
	return &OrderReader{
		scanner: bufio.NewScanner(r),
	}
}
