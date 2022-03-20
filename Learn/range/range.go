package main

import "fmt"

func main(){
	var nums = [] int {2,3,4}
	var sum int = 0
	var num int

	for _, num = range nums{
		sum = sum + num
	}

	fmt.Printf("Sum: %d", sum)
}
