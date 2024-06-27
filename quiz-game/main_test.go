package main

import (
	"fmt"
	"log"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseRow(t *testing.T) {
	expected := []problem{
		{"one", "two"},
	}
	rows := [][]string{
		{"one", "two"},
	}

	output, _ := parseRows(rows)

	// Test return type
	if reflect.TypeOf(output) != reflect.TypeOf(expected) {
		log.Fatalf("parseRows(%v) output type: %T != expected type: %T", rows, output, expected)
	}

	// Test size of return slice
	if len(output) != len(expected) {
		log.Fatalf("parseRows(%v) output length %v does not equal expected length: %v", len(output), rows, len(expected))
	}

	// Test equality assert on return value rather than index comparison
	if !assert.Equal(t, output, expected) {
		log.Fatalf("parseRows(%v) output value: %v not equal to expected value: %v", rows, output[0], expected[0])
	}

	// Test error handing with invalid rows
	invalidRows := [][]string{
		{"one", "two", "three"},
	}
	_, err := parseRows(invalidRows)
	expectedErr := fmt.Errorf("unexpected element content, expected question and answer, got: [one two three]")

	if !assert.Equal(t, err, expectedErr) {
		log.Fatalf("returned error:\n%v\ndid not equal expected:\n%v", err, expectedErr)
	}
}

func TestFileReader(t *testing.T) {
	expected := [][]string{
		{"one", "two"},
		{"one", "two"},
		{"one", "two"},
	}

	testFile := "test-data.csv"
	output, _ := fileReader(testFile)

	//Test return type
	if reflect.TypeOf(output) != reflect.TypeOf(expected) {
		log.Fatalf("parseRows(%v) output type: %T != expected type: %T", testFile, output, expected)
	}

	// Test size of return 2D slice
	if len(output) != len(expected) {
		log.Fatalf("fileReader(%v) output length: %v does not equal expected length: %v", testFile, len(output), len(expected))
	}

	// Test equality assert on return value rather than index comparison
	if !assert.Equal(t, output, expected) {
		log.Fatalf("parseRows(%v) output value: %v not equal to expected value: %v", testFile, output[0], expected[0])
	}

	// Test error handing with invalid filepath
	_, err := fileReader("invalidFilepath")
	expectedErr := fmt.Errorf("unable to open file from invalidFilepath, open invalidFilepath: no such file or directory")

	if !assert.Equal(t, err, expectedErr) {
		log.Fatalf("returned error:\n%v\ndid not equal expected:\n%v", err, expectedErr)
	}
}
