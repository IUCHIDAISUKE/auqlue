package main

import "fmt"

// Variable is a collection of variables
type Variable struct {
	data int
}

func main() {
	var x Variable
	x.data = 10
	fmt.Println(x.data)
}
