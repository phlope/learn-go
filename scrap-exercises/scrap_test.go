package main

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertToMap(t *testing.T) {

	fruitSlice := []string{"pineapple", "orange", "banana", "kiwi"}
	output := convertToMap(fruitSlice)

	expected := map[string]float64{
		"pineapple": 25.00,
		"orange":    25.00,
		"banana":    25.00,
		"kiwi":      25.00,
	}

	if !assert.Equal(t, output, expected) {
		log.Fatalf("Returned value: %[1]v (type %[1]T) does not match expected: %[2]v (type %[2]T)", output, expected)
	}

	fruitSlice = []string{"pineapple"}
	output = convertToMap(fruitSlice)

	expected = map[string]float64{
		"pineapple": 100.00,
	}

	if !assert.Equal(t, output, expected) {
		log.Fatalf("Returned value: %[1]v (type %[1]T) does not match expected: %[2]v (type %[2]T)", output, expected)
	}
}

func TestCalculateTotal(t *testing.T) {

	cart := []cartItem{
		{"apple", 2.99, 3},
		{"orange", 1.50, 8},
		{"banana", .49, 12},
	}

	output := calculateTotal(cart)
	expected := 26.85

	if !assert.Equal(t, output, expected) {
		log.Fatalf("Returned value: %[1]v (type %[1]T) does not match expected: %[2]v (type %[2]T)", output, expected)
	}

}
