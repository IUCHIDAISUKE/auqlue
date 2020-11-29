package main

import "fmt"

// Variable is a collection of variables
type Variable struct {
	data float64
}

// Function is a collection of functions
type Function struct {
	procedure func(n float64) float64
}

func (f Function) run(v Variable) float64 {
	res := v.data * v.data
	return res
}

func main() {
	var x Variable
	var f Function
	// fmt.Println(x.data)
	// fmt.Println(f.procedure)
	x.data = 10
	y := f.run(x)
	fmt.Println(y)
}
