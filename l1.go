package main

import "math/rand"
import "fmt"

func main1() {
	/*//you need a variable there are 2 ways:
	var x_1 = 10
	var x2 int
	x2 = 10
	var x3 int = 10
	// we can also use
	x4 := 10 //in this situation we must initialise it*/

	//arrays
	var x = []int{1, 1}
	fmt.Println(cap(x))
	for k := 0; k < len(x); k = k + 1 {
		j := rand.Int()
		if x[k] != 1 {
			x[k] = j % 10
		}
	}
	fmt.Println(x)

	//Slices
	z := make([]int, 5, 10) // Slice di 5 elementi inizializzati a 0
	initialize(z)
	fmt.Println(z)
}

func initialize(x []int) {
	for k := 0; k < len(x); k++ {
		x[k] = rand.Int()
	}
}
