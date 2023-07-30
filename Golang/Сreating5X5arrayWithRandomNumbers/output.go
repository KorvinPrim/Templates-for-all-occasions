package main

import (
	"fmt"
	"github.com/zs5460/art"
)

func WordGreeting() error {
	_, err := fmt.Println(art.String("Rand array"))
	if err != nil {
		return err
	}
	return nil
}

func Print(intArray *IntArray) error {
	arr, err := intArray.Get()
	if err != nil {
		return err
	}

	for rowIndex, rows := range arr {
		fmt.Printf("%d. %v\n", rowIndex+1, rows)
	}

	return nil
}
