package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type GetUniRandIntInp struct {
	SizeX       int
	SizeY       int
	RandomLimit int
}

func (i GetUniRandIntInp) Validate() error {
	if i.SizeX <= 0 || i.SizeY <= 0 {
		return errors.New("array length can't negetive or 0")
	}
	if i.RandomLimit <= 0 {
		return errors.New("random limit can't be 0")
	}

	if i.RandomLimit < i.SizeX*i.SizeY {
		return errors.New(
			fmt.Sprintf("Can't generate unique numbers with limit %v for %vx%v size aray", i.RandomLimit, i.SizeX, i.SizeY))
	}

	return nil
}

type Randomizer struct {
	userIntegers map[int]bool
}

func NewRandomizer() *Randomizer {
	rand.Seed(time.Now().Unix())
	return &Randomizer{
		userIntegers: make(map[int]bool),
	}
}
