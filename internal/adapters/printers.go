package adapters

import (
	"fmt"
	"io"
	"sort"
	"strconv"

	"github.com/martishin/load-balancer/internal/domain"
	"github.com/martishin/load-balancer/internal/models"
)

func formatWithCommas(n int32, width int) string {
	sign := ""
	if n < 0 {
		sign = "-"
		n = -n
	}
	s := strconv.Itoa(int(n))
	for i := len(s) - 3; i > 0; i -= 3 {
		s = s[:i] + "," + s[i:]
	}
	return fmt.Sprintf("%*s", width, sign+s)
}

func PrintBook(writer io.Writer, book domain.Book) {
	buyOrders := book.BuyOrders()
	sellOrders := book.SellOrders()

	sort.SliceStable(buyOrders, func(i, j int) bool {
		if buyOrders[i].Price == buyOrders[j].Price {
			return buyOrders[i].ArrivalOrder < buyOrders[j].ArrivalOrder
		}
		return buyOrders[i].Price > buyOrders[j].Price
	})

	sort.SliceStable(sellOrders, func(i, j int) bool {
		if sellOrders[i].Price == sellOrders[j].Price {
			return sellOrders[i].ArrivalOrder < sellOrders[j].ArrivalOrder
		}
		return sellOrders[i].Price < sellOrders[j].Price
	})

	maxRows := max(len(buyOrders), len(sellOrders))

	for i := 0; i < maxRows; i++ {
		var buyQuantity, buyPrice, sellPrice, sellQuantity string

		if i < len(buyOrders) {
			buyQuantity = formatWithCommas(buyOrders[i].Quantity, 9) // Right-align with width 9
			buyPrice = formatWithCommas(buyOrders[i].Price, 6)       // Right-align with width 6
		} else {
			buyQuantity = "         " // 9 spaces
			buyPrice = "      "       // 6 spaces
		}

		if i < len(sellOrders) {
			sellPrice = formatWithCommas(sellOrders[i].Price, 6)       // Right-align with width 6
			sellQuantity = formatWithCommas(sellOrders[i].Quantity, 9) // Right-align with width 9
		} else {
			sellPrice = "      "       // 6 spaces
			sellQuantity = "         " // 9 spaces
		}

		fmt.Fprintf(writer, " %s %s | %s %s\n", buyQuantity, buyPrice, sellPrice, sellQuantity)
	}
}

func PrintMatch(writer io.Writer, match *models.Match) {
	fmt.Fprintf(writer, "trade %d,%d,%d,%d\n",
		match.AggressionOrderID,
		match.RestingOrderID,
		match.Price,
		match.Quantity,
	)
}

func PrintOrder(writer io.Writer, order *models.Order) {
	fmt.Fprintf(writer, "%d,%s,%d,%d\n",
		order.ID,
		order.Type,
		order.Price,
		order.Quantity,
	)
}
