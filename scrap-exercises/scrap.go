package main

import (
	"fmt"
	"math"
)

func main() {
	// 1: Slice to Map
	fruitSlice := []string{"apple", "orange", "banana"}
	fruitMap := convertToMap(fruitSlice)
	fmt.Print(fruitMap)

	// 2: Cart total
	var cart []cartItem
	var apples = cartItem{"apple", 2.99, 3}
	var oranges = cartItem{"orange", 1.50, 8}
	var bananas = cartItem{"banana", .49, 12}
	cart = append(cart, apples, oranges, bananas)
	result := calculateTotal(cart)
	fmt.Print(result)
}

// Convert Slice to Map ensuring values of fruit always sum to 100
func convertToMap(items []string) map[string]float64 {
	result := make(map[string]float64)
	for _, fruit := range items {
		value := float64(100 / len(items))
		result[fruit] = value
	}
	return result
}

type cartItem struct {
	name     string
	price    float64
	quantity int
}

// Calculate cart price
func calculateTotal(cart []cartItem) float64 {
	total := 0.0
	for _, item := range cart {
		total += float64(item.quantity) * item.price
	}
	return math.Round((total)*100) / 100
}
