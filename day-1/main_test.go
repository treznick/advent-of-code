package main

import "testing"

func TestGetMass(t *testing.T) {
	tables := []struct {
		input    int64
		expected int64
	}{
		{100, 31},
		{15, 3},
	}

	for _, table := range tables {
		mass := getMass(table.input)
		if mass != table.expected {
			t.Errorf("expeted '%d' but got '%d'", table.expected, mass)
		}
	}
}
