package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func runProgram(input string) []int64 {
	program := toIntegerSlice(input)
	programSize := 4

	for i := 0; i < len(program); i += programSize {
		j := i + programSize
		if j > len(program) {
			j = len(program)
		}

		intCode := program[i:j]

		if intCode[0] == 99 {
			break
		}

		runIntCode(&program, intCode)
	}

	return program
}

func runIntCode(program *[]int64, operation []int64) {
	switch operation[0] {
	case 1:
		(*program)[operation[3]] = (*program)[operation[1]] + (*program)[operation[2]]
	case 2:
		(*program)[operation[3]] = (*program)[operation[1]] * (*program)[operation[2]]
	}
}

func repairProgram(input string, noun int64, verb int64) string {
	arr := strings.Split(input, ",")
	arr[1] = strconv.FormatInt(noun, 10)
	arr[2] = strconv.FormatInt(verb, 10)
	return strings.Join(arr, ",")
}

func fromIntegerSlice(input []int64) string {
	output := make([]string, len(input))

	for index, value := range input {
		output[index] = strconv.FormatInt(int64(value), 10)
	}

	return strings.Join(output, ",")
}

func toIntegerSlice(input string) []int64 {
	arr := strings.Split(input, ",")
	newSlice := make([]int64, len(arr))

	for i, c := range arr {
		stripped := strings.TrimSpace(c)
		parsed, err := strconv.ParseInt(stripped, 10, 64)
		if err != nil {
			panic(err)
		}
		newSlice[i] = int64(parsed)
	}

	return newSlice
}

func loadProgram(path string) string {
	contents, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	return string(contents)
}

func main() {
	path := "./input.txt"

out:
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			s := loadProgram(path)

			repaired := repairProgram(s, int64(noun), int64(verb))
			output := runProgram(repaired)
			if output[0] == 19690720 {
				fmt.Println("the answer is", 100*noun+verb)
				break out
			}
		}
	}
}
