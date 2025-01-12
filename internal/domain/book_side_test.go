package domain_test

import (
	"reflect"
	"testing"

	"github.com/martishin/load-balancer/internal/domain"
	"github.com/martishin/load-balancer/internal/models"
)

func TestBookSide(t *testing.T) {
	orderLowPrice := &models.Order{
		ID:       10005,
		Type:     models.BuyOrder,
		Price:    100,
		Quantity: 20000,
	}
	orderHighPrice := &models.Order{
		ID:       10006,
		Type:     models.BuyOrder,
		Price:    105,
		Quantity: 10000,
	}

	t.Run("buy book side", func(t *testing.T) {
		bookSide := domain.NewBookSide(models.BuyOrder)
		bookSide.Push(orderLowPrice)
		bookSide.Push(orderHighPrice)

		want := []*models.Order{orderHighPrice, orderLowPrice}

		var got []*models.Order
		got = append(got, bookSide.Pop())
		got = append(got, bookSide.Pop())

		if len(got) != 2 {
			t.Errorf("got %d orders, want %d", len(got), 2)
		}

		if !reflect.DeepEqual(want, got) {
			t.Fatalf("wanted %v, got %v", want, got)
		}
	})

	t.Run("sell book side", func(t *testing.T) {
		bookSide := domain.NewBookSide(models.SellOrder)
		bookSide.Push(orderLowPrice)
		bookSide.Push(orderHighPrice)

		want := []*models.Order{orderLowPrice, orderHighPrice}

		var got []*models.Order
		got = append(got, bookSide.Pop())
		got = append(got, bookSide.Pop())

		if len(got) != 2 {
			t.Errorf("got %d orders, want %d", len(got), 2)
		}

		if !reflect.DeepEqual(want, got) {
			t.Fatalf("wanted %v, got %v", want, got)
		}
	})

	t.Run("order shouldn't matter", func(t *testing.T) {
		bookSide := domain.NewBookSide(models.SellOrder)
		bookSide.Push(orderLowPrice)
		bookSide.Push(orderHighPrice)

		want := []*models.Order{orderLowPrice, orderHighPrice}

		var got []*models.Order
		got = append(got, bookSide.Pop())
		got = append(got, bookSide.Pop())

		if !reflect.DeepEqual(want, got) {
			t.Fatalf("wanted %v, got %v", want, got)
		}

		bookSide.Push(orderHighPrice)
		bookSide.Push(orderLowPrice)
		got = nil
		got = append(got, bookSide.Pop())
		got = append(got, bookSide.Pop())
		if !reflect.DeepEqual(want, got) {
			t.Fatalf("wanted %v, got %v", want, got)
		}
	})
}
