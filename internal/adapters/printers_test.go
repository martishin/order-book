package adapters_test

import (
	"strings"
	"testing"

	"github.com/martishin/load-balancer/internal/adapters"
	"github.com/martishin/load-balancer/internal/domain"
	"github.com/martishin/load-balancer/internal/models"
)

func TestPrintBook(t *testing.T) {
	t.Run("should print book", func(t *testing.T) {
		book := domain.NewInMemoryBook()

		book.AddOrder(&models.Order{
			ID:       10000,
			Type:     models.BuyOrder,
			Price:    99,
			Quantity: 50000,
		})
		book.AddOrder(&models.Order{
			ID:       10001,
			Type:     models.SellOrder,
			Price:    100,
			Quantity: 500,
		})

		expected := "    50,000     99 |    100       500\n"

		var buf strings.Builder
		adapters.PrintBook(&buf, book)
		got := buf.String()

		if got != expected {
			t.Errorf("unexpected output:\nGot:\n%s\nWant:\n%s", buf.String(), expected)
		}
	})

	t.Run("should print book in correct order", func(t *testing.T) {
		book := domain.NewInMemoryBook()

		book.AddOrder(&models.Order{
			ID:       10000,
			Type:     models.BuyOrder,
			Price:    99,
			Quantity: 50000,
		})
		book.AddOrder(&models.Order{
			ID:       10001,
			Type:     models.SellOrder,
			Price:    100,
			Quantity: 500,
		})
		book.AddOrder(&models.Order{
			ID:       10002,
			Type:     models.SellOrder,
			Price:    100,
			Quantity: 10000,
		})
		book.AddOrder(&models.Order{
			ID:       10003,
			Type:     models.BuyOrder,
			Price:    99,
			Quantity: 500,
		})

		expected := "    50,000     99 |    100       500\n" +
			"       500     99 |    100    10,000\n"

		var buf strings.Builder
		adapters.PrintBook(&buf, book)
		got := buf.String()

		if got != expected {
			t.Errorf("unexpected output:\nGot:\n%s\nWant:\n%s", got, expected)
		}
	})
}

func TestPrintMatch(t *testing.T) {
	match := &models.Match{
		AggressionOrderID: 10006,
		RestingOrderID:    10001,
		Price:             100,
		Quantity:          500,
	}

	var buf strings.Builder
	adapters.PrintMatch(&buf, match)

	expected := "trade 10006,10001,100,500\n"
	if buf.String() != expected {
		t.Errorf("unexpected output: got %q, want %q", buf.String(), expected)
	}
}

func TestPrintOrder(t *testing.T) {
	order := &models.Order{
		ID:       10000,
		Type:     models.BuyOrder,
		Price:    99,
		Quantity: 50000,
	}

	var buf strings.Builder
	adapters.PrintOrder(&buf, order)

	expected := "10000,B,99,50000\n"
	if buf.String() != expected {
		t.Errorf("unexpected output: got %q, want %q", buf.String(), expected)
	}
}
