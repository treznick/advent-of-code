package main

import (
	"reflect"
	"testing"
)

func TestRunProgram(t *testing.T) {
	tables := []struct {
		input    string
		expected []int64
	}{
		{"1,0,0,0,99", []int64{2, 0, 0, 0, 99}},
		{"2,3,0,3,99", []int64{2, 3, 0, 6, 99}},
		{"2,4,4,5,99,0", []int64{2, 4, 4, 5, 99, 9801}},
		{"1,1,1,4,99,5,6,0,99", []int64{30, 1, 1, 4, 2, 5, 6, 0, 99}},
		{"1,9,10,3,2,3,11,0,99,30,40,50", []int64{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50}},
	}

	for _, table := range tables {
		output := runProgram(table.input)
		if !reflect.DeepEqual(table.expected, output) {
			t.Errorf("expeted '%q' but got '%q'", table.expected, output)
		}
	}
}

func TestRepairProgram(t *testing.T) {
	tables := []struct {
		input    string
		noun     int64
		verb     int64
		expected string
	}{
		{"1,0,0,0,99", 12, 2, "1,12,2,0,99"},
		{"1,0,0,0,99", 1, 2, "1,1,2,0,99"},
	}

	for _, table := range tables {
		output := repairProgram(table.input, table.noun, table.verb)
		if output != table.expected {
			t.Errorf("expeted '%s' but got '%s'", table.expected, output)
		}
	}
}

func TestToIntegerSlice(t *testing.T) {
	tables := []struct {
		input    string
		expected []int64
	}{
		{"1,2,3,4,5", []int64{1, 2, 3, 4, 5}},
		{"1,2,3\n", []int64{1, 2, 3}},
	}

	for _, table := range tables {
		output := toIntegerSlice(table.input)
		if !reflect.DeepEqual(table.expected, output) {
			t.Errorf("expeted '%q' but got '%q'", table.expected, output)
		}
	}
}

func TestFrom(t *testing.T) {
	input := []int64{1, 2, 3, 4, 5}
	expected := "1,2,3,4,5"

	output := fromIntegerSlice(input)

	if output != expected {
		t.Errorf("expeted '%s' but got '%s'", expected, output)
	}
}
