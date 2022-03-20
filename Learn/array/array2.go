package main

import "fmt"

func main() {
	var i int
	balance := [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	for i = 0; i < 5; i++ {
		fmt.Printf("Balance[%d] = %f\n", i, balance[i])
	}

	balance2 := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	for i = 0; i < 5; i++ {
		fmt.Printf("Balance2[%d] = %f\n", i, balance2[i])
	}

	balance3 := [5]float32{1: 2.0, 3: 7.0}
	for i = 0; i < 5; i++ {
		fmt.Printf("Balance3[%d] = %f\n", i, balance3[i])
	}

}
