package main

import "testing"

func TestSum(t *testing.T) {
	result := sum(2, 2)
	expected := 4
	if result != expected {
		t.Errorf("sum(2, 2) = %d; want %d", result, expected)
	}
}

func TestSumNegative(t *testing.T) {
	result := sum(-2, -3)
	expected := -5
	if result != expected {
		t.Errorf("sum(-2, -3) = %d; want %d", result, expected)
	}
}

func TestSumZero(t *testing.T) {
	result := sum(0, 0)
	expected := 0
	if result != expected {
		t.Errorf("sum(0, 0) = %d; want %d", result, expected)
	}
}

func TestSumMixedSigns(t *testing.T) {
	result := sum(-5, 3)
	expected := -2
	if result != expected {
		t.Errorf("sum(-5, 3) = %d; want %d", result, expected)
	}
}
