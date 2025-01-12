package adapters_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/martishin/load-balancer/internal/adapters"
	"github.com/martishin/load-balancer/internal/models"
)

func TestReader(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected *models.Order
	}{
		{
			name:  "buy order",
			input: "10000,B,98,25500",
			expected: &models.Order{
				ID:       10000,
				Type:     models.BuyOrder,
				Price:    98,
				Quantity: 25500,
			},
		},
		{
			name:  "sell order",
			input: "10005,S,105,20000",
			expected: &models.Order{
				ID:       10005,
				Type:     models.SellOrder,
				Price:    105,
				Quantity: 20000,
			},
		},
	}

	for _, c := range cases {
		t.Run(
			c.name, func(t *testing.T) {
				orderReader := adapters.NewOrderReader(strings.NewReader(c.input))
				got, err := orderReader.Next()

				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}

				if !reflect.DeepEqual(got, c.expected) {
					t.Errorf("expected: %+v, got: %+v", c.expected, got)
				}
			},
		)
	}
}
