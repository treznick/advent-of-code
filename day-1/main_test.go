package main

import "testing"

func TestGetMass(t *testing.T) {
	tables := []struct {
		input    int32
		expected int32
	}{
		{100, 31},
		{15, 3},
		{2, 0},
		{6, 0},
	}

	for _, table := range tables {
		mass := getMass(table.input)
		if mass != table.expected {
			t.Errorf("expeted '%d' but got '%d'", table.expected, mass)
		}
	}
}

func TestCalculateFuelMass(t *testing.T) {
	tables := []struct {
		input    int32
		expected int32
	}{
		{2, 0},
		{6, 0},
		{14, 2},
		{1969, 966},
		{100756, 50346},
	}

	for _, table := range tables {
		mass := calculateFuelMass(table.input)
		if mass != table.expected {
			t.Errorf("expeted '%d' but got '%d'", table.expected, mass)
		}
	}
}
