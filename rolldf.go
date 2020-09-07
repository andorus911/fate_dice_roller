package main

import (
	"fmt"
	"math/rand"
)

func main() {
	r := Roll{}
	r.Make()
	fmt.Println(r)
}

type Roll struct {
	Dice [4]int
	Sum  int
}

func (r *Roll) Make() {
	r.Sum = 0
	for i := 0; i < 4; i++ {
		r.Dice[i] = rand.Intn(3) - 1
		r.Sum += r.Dice[i]
	}
}
