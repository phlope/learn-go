package main

import (
	"log"
	"testing"
)

func TestExtractDigitPt1(t *testing.T) {
	charA := '3'
	charB := 'p'

	extracted, isNum := extractDigitPt1(byte(charA))

	if extracted != 3 || isNum != true {
		log.Fatalf("Expected 10, true, found %v, %v", extracted, isNum)
	}

	extracted, isNum = extractDigitPt1(byte(charB))

	if extracted != 0 || isNum != false {
		log.Fatalf("Expected 0, false, found %v, %v", extracted, isNum)
	}
}
