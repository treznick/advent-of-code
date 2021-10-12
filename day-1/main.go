package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func getMass(m int64) int64 {
	return int64(math.Floor(float64(m)/3)) - 2
}

func main() {
	sum := int64(0)
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		i, err := strconv.ParseInt(scanner.Text(), 10, 64)

		if err != nil {
			panic(err)
		}

		sum += getMass(i)
	}

	fmt.Println(sum)
}
