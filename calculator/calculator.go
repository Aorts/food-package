package calculator

import "fmt"

var OrderCount = map[string]int{
	"red":    0,
	"green":  0,
	"blue":   0,
	"yellow": 0,
	"pink":   0,
	"purple": 0,
	"orange": 0,
}

var Prices = map[string]float64{
	"red":    50,
	"green":  40,
	"blue":   30,
	"yellow": 50,
	"pink":   80,
	"purple": 90,
	"orange": 120,
}

var bundleItems = map[string]bool{
	"green":  true,
	"pink":   true,
	"orange": true,
}

func OrderFood(orderName string, quantity int) {
	if _, exists := Prices[orderName]; !exists {
		println("Invalid order name")
		return
	}
	OrderCount[orderName] += quantity
	fmt.Println("All items in the order are: ")
	for k, v := range OrderCount {
		fmt.Println(k, " : ", v)
	}
}

func CheckOut(isMember bool) float64 {
	var totalPrice float64
	for k, v := range OrderCount {
		if v > 0 {
			itemPrice := Prices[k] * float64(v)
			if bundleItems[k] && v >= 2 {
				bundles := v / 2
				itemPrice -= Prices[k] * 0.05 * float64(bundles*2)
			}
			totalPrice += itemPrice
		}
	}
	if isMember {
		totalPrice *= 0.9
	}
	fmt.Println("Total price: ", totalPrice)
	return totalPrice
}
