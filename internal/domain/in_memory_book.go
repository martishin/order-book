package domain

import (
	"github.com/martishin/load-balancer/internal/models"
)

type InMemoryBook struct {
	buyOrders    *BookSide
	sellOrders   *BookSide
	orderCounter int32
}

func (b *InMemoryBook) BuyOrders() []*models.Order {
	return b.buyOrders.Orders()
}

func (b *InMemoryBook) SellOrders() []*models.Order {
	return b.sellOrders.Orders()
}

func (b *InMemoryBook) AddOrder(order *models.Order) []*models.Match {
	b.orderCounter++
	order.ArrivalOrder = b.orderCounter

	var orderSide *BookSide
	var otherSide *BookSide
	if order.Type == models.BuyOrder {
		orderSide = b.buyOrders
		otherSide = b.sellOrders
	} else {
		orderSide = b.sellOrders
		otherSide = b.buyOrders
	}

	matches := b.match(order, otherSide)
	if order.Quantity > 0 {
		orderSide.Push(order)
	}
	return matches
}

func (b *InMemoryBook) match(order *models.Order, otherSide *BookSide) []*models.Match {
	var matches []*models.Match
	for order.Quantity > 0 && otherSide.Len() > 0 {
		bestMatch := otherSide.Pop()

		if (order.Type == models.BuyOrder && order.Price >= bestMatch.Price) ||
			(order.Type == models.SellOrder && order.Price <= bestMatch.Price) {
			matchQuantity := min(order.Quantity, bestMatch.Quantity)
			matchPrice := min(order.Price, bestMatch.Price)
			order.Quantity -= matchQuantity
			bestMatch.Quantity -= matchQuantity

			matches = append(
				matches, &models.Match{
					AggressionOrderID: order.ID,
					RestingOrderID:    bestMatch.ID,
					Price:             matchPrice,
					Quantity:          matchQuantity,
				},
			)

			if bestMatch.Quantity > 0 {
				otherSide.Push(bestMatch)
			}
		} else {
			otherSide.Push(bestMatch)
			break
		}
	}
	return matches
}

func NewInMemoryBook() *InMemoryBook {
	return &InMemoryBook{
		buyOrders:  NewBookSide(models.BuyOrder),
		sellOrders: NewBookSide(models.SellOrder),
	}
}
