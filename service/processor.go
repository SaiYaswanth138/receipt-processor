package service

import (
	"math"
	"strconv"
	"strings"
	"time"
	"unicode"

	"receipt-processor/model"
)

func CalculatePoints(receipt model.Receipt) int {
	totalPoints := 0

	for _, char := range receipt.Retailer {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			totalPoints = totalPoints + 1
		}
	}

	totalAmount, _ := strconv.ParseFloat(receipt.Total, 64)

	if totalAmount == float64(int(totalAmount)) {
		totalPoints = totalPoints + 50
	}

	if math.Mod(totalAmount, 0.25) == 0 {
		totalPoints = totalPoints + 25
	}

	totalPoints = totalPoints + (len(receipt.Items) / 2) * 5


	for _, item := range receipt.Items {
		desc := strings.TrimSpace(item.ShortDescription)
		if len(desc)%3 == 0 {
			itemPrice, _ := strconv.ParseFloat(item.Price, 64)
			bonus := math.Ceil(itemPrice * 0.2)
			totalPoints = totalPoints + int(bonus)
		}
	}

	purchaseDate, _ := time.Parse("2006-01-02", receipt.PurchaseDate)
	if purchaseDate.Day()%2 != 0 {
		totalPoints = totalPoints + 6
	}

	purchaseTime, _ := time.Parse("15:04", receipt.PurchaseTime)
	if purchaseTime.Hour() == 14 {
		totalPoints = totalPoints + 10
	}

	return totalPoints
}
