package domain

import "github.com/martishin/load-balancer/internal/models"

type Book interface {
	AddOrder(order *models.Order) []*models.Match
	BuyOrders() []*models.Order
	SellOrders() []*models.Order
}
