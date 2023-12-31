package main

import (
	"errors"
)

type IntArray struct {
	sizeX int
	sizeY int
	arr   [][]int
}

func NewIntArray(sizeX int, sizeY int) (*IntArray, error) {
	if sizeX <= 0 || sizeY <= 0 {
		return nil, errors.New("array size can't be negative or 0")
	}
	return &IntArray{sizeX: sizeX, sizeY: sizeY}, nil
}

func (a *IntArray) Fill(valuses []int) error {
	if len(valuses) < a.sizeY*a.sizeY {
		return errors.New("valuses array has less element than target array")
	}
	rows := make([][]int, a.sizeX)
	counter := 0

	for row := 0; row < a.sizeX; row++ {
		columns := make([]int, a.sizeY)
		rows[row] = columns

		for col := 0; col < a.sizeY; col++ {
			rows[row][col] = valuses[counter]
			counter++
		}
	}
	a.arr = rows
	return nil
}

func (a *IntArray) Get() ([][]int, error) {
	if a.arr == nil {
		return nil, errors.New("array is not filled with values")
	}
	return a.arr, nil
}
