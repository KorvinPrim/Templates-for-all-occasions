package main

import (
	"log"
)

const (
	rowsCount    = 5
	columnsCount = 5
	randomLimit  = 100
)

func main() {
	err := WordGreeting()
	if err != nil {
		log.Fatal(err)
	}

	randomizer := NewRandomizer()
	randomIntegers, err := randomizer.GetUniqueRandomIntegers(GetUniRandIntInp{
		SizeX:       rowsCount,
		SizeY:       columnsCount,
		RandomLimit: randomLimit,
	})
	if err != nil {
		log.Fatal(err)
	}

	array, err := NewIntArray(rowsCount, columnsCount)
	if err != nil {
		log.Fatal(err)
	}

	if err := array.Fill(randomIntegers); err != nil {
		log.Fatal(err)
	}

	if err := Print(array); err != nil {
		log.Fatal(err)
	}
}
