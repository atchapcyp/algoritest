package main

import (
	"fmt"
	"testing"
)

func Test_atoi_5001_Must_Be_5001_int(t *testing.T) {
	expected := 5001
	actual, _ := ezatoi("5001", 0, 4)
	if actual != expected {
		fmt.Printf("Expected : %v  but get %v", expected, actual)
	}
}
