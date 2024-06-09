package main

import (
	"log"
	"reflect"
	"testing"
)

func TestParseRow(t *testing.T) {
	expected := []problem{
		{"one", "two"},
	}
	rows := [][]string{
		{"one", "two"},
	}

	output := parseRows(rows)

	// Test return type
	if reflect.TypeOf(output) != reflect.TypeOf(expected) {
		log.Fatalf("parseRows(%v) output type: %T != expected type: %T", rows, output, expected)
	}

	// Test size of return slice
	if len(output) != len(expected) {
		log.Fatalf("parseRows(%v) output length %v does not equal expected length: %v", len(output), rows, len(expected))
	}

	// Test value of return slice element
	if output[0] != expected[0] {
		log.Fatalf("parseRows(%v) output value: %v not equal to expected value: %v", rows, output[0], expected[0])
	}
}

func TestFileReader(t *testing.T) {
	expected := [][]string{
		{"one", "two"},
		{"one", "two"},
		{"one", "two"},
	}

	testFile := "test-data.csv"
	output := fileReader(testFile)

	//Test return type
	if reflect.TypeOf(output) != reflect.TypeOf(expected) {
		log.Fatalf("parseRows(%v) output type: %T != expected type: %T", testFile, output, expected)
	}

	// Test size of return 2D slice
	if len(output) != len(expected) {
		log.Fatalf("fileReader(%v) output length: %v does not equal expected length: %v", testFile, len(output), len(expected))
	}

	//Compare each slice element
	for i := range output {
		for j := range output[i] {
			if output[i][j] != expected[i][j] {
				log.Fatalf("parseRows(%v) output value: %v not equal to expected value: %v", testFile, output[i][j], expected[i][j])
			}
		}
	}
}
