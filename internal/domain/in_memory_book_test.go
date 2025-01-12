package domain_test

import (
	"reflect"
	"testing"

	"github.com/martishin/load-balancer/internal/domain"
	"github.com/martishin/load-balancer/internal/models"
)

func TestInMemoryBook(t *testing.T) {
	t.Run("match buy", func(t *testing.T) {
		book := domain.NewInMemoryBook()
		book.AddOrder(&models.Order{
			ID:       10000,
			Type:     models.BuyOrder,
			Price:    110,
			Quantity: 25500,
		})
		want := []*models.Match{&models.Match{
			AggressionOrderID: 10005,
			RestingOrderID:    10000,
			Price:             105,
			Quantity:          20000,
		}}
		got := book.AddOrder(&models.Order{
			ID:       10005,
			Type:     models.SellOrder,
			Price:    105,
			Quantity: 20000,
		})

		if !reflect.DeepEqual(want, got) {
			t.Errorf("want: %v, got: %v", want[0], got[0])
		}
	})

	t.Run("match sell", func(t *testing.T) {
		book := domain.NewInMemoryBook()
		book.AddOrder(&models.Order{
			ID:       10005,
			Type:     models.SellOrder,
			Price:    105,
			Quantity: 20000,
		})
		want := []*models.Match{&models.Match{
			AggressionOrderID: 10000,
			RestingOrderID:    10005,
			Price:             105,
			Quantity:          20000,
		}}
		got := book.AddOrder(&models.Order{
			ID:       10000,
			Type:     models.BuyOrder,
			Price:    110,
			Quantity: 25500,
		})

		if !reflect.DeepEqual(want, got) {
			t.Errorf("want: %v, got: %v", want[0], got[0])
		}
	})

	t.Run("can't match buy", func(t *testing.T) {
		book := domain.NewInMemoryBook()
		book.AddOrder(&models.Order{
			ID:       10000,
			Type:     models.BuyOrder,
			Price:    100,
			Quantity: 25500,
		})
		var want []*models.Match
		got := book.AddOrder(&models.Order{
			ID:       10005,
			Type:     models.SellOrder,
			Price:    105,
			Quantity: 20000,
		})

		if !reflect.DeepEqual(want, got) {
			t.Errorf("want: %v, got: %v", want, got)
		}
	})

	t.Run("can't match sell", func(t *testing.T) {
		book := domain.NewInMemoryBook()
		book.AddOrder(&models.Order{
			ID:       10005,
			Type:     models.SellOrder,
			Price:    105,
			Quantity: 20000,
		})
		var want []*models.Match
		got := book.AddOrder(&models.Order{
			ID:       10000,
			Type:     models.BuyOrder,
			Price:    100,
			Quantity: 25500,
		})

		if !reflect.DeepEqual(want, got) {
			t.Errorf("want: %v, got: %v", want, got)
		}
	})
}
