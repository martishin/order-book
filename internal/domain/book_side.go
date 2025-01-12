package domain

import (
	"container/heap"

	"github.com/martishin/load-balancer/internal/models"
)

type BookSide struct {
	ordersQueue OrdersQueue
}

func (bs *BookSide) Push(order *models.Order) {
	heap.Push(&bs.ordersQueue, order)
}

func (bs *BookSide) Pop() *models.Order {
	if bs.ordersQueue.Len() == 0 {
		return nil
	}
	return heap.Pop(&bs.ordersQueue).(*models.Order)
}

func (bs *BookSide) Len() int {
	return bs.ordersQueue.Len()
}

func (bs *BookSide) Orders() []*models.Order {
	return bs.ordersQueue.orders
}

func NewBookSide(orderType models.OrderType) *BookSide {
	var cmp func(lhs, rhs *models.Order) bool
	if orderType == models.SellOrder {
		cmp = func(lhs, rhs *models.Order) bool { return lhs.Price < rhs.Price }
	} else {
		cmp = func(lhs, rhs *models.Order) bool { return lhs.Price > rhs.Price }
	}

	return &BookSide{
		ordersQueue: OrdersQueue{
			cmp: cmp,
		},
	}
}

type OrdersQueue struct {
	orders []*models.Order
	cmp    func(lhs, rhs *models.Order) bool
}

func (pq OrdersQueue) Len() int {
	return len(pq.orders)
}

func (pq OrdersQueue) Less(i, j int) bool {
	return pq.cmp(pq.orders[i], pq.orders[j])
}

func (pq OrdersQueue) Swap(i, j int) {
	pq.orders[i], pq.orders[j] = pq.orders[j], pq.orders[i]
}

func (pq *OrdersQueue) Push(x any) {
	pq.orders = append(pq.orders, x.(*models.Order))
}

func (pq *OrdersQueue) Pop() interface{} {
	n := len(pq.orders)
	item := pq.orders[n-1]
	pq.orders = pq.orders[:n-1]
	return item
}
