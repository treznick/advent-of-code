package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func getMass(m int32) int32 {
	return int32(math.Max((math.Floor(float64(m)/3) - 2), float64(0)))
}

type module struct {
	mass     int32
	fuelMass int32
}

func (m *module) totalMass() int32 {
	return m.mass + m.fuelMass
}

func newModule(m int32) *module {
	p := module{mass: m, fuelMass: calculateFuelMass(m)}
	return &p
}

func calculateFuelMass(m int32) int32 {
	i := m
	agg := int32(0)
	for getMass(i) > 0 {
		i = getMass(i)
		agg += i
	}

	return agg
}

func main() {
	sum := int32(0)
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		i, err := strconv.ParseInt(scanner.Text(), 10, 32)

		if err != nil {
			panic(err)
		}

		m := newModule(int32(i))
		sum += m.fuelMass
	}

	fmt.Println(sum)
}
