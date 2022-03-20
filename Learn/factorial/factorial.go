package main

import "fmt"

func main(){
	var i int = 15

	fmt.Printf("%d 's factorial is %d\n", i, factorial(i))
}

func factorial(n int) int{
	var result int
	if n > 0{
		result = n * factorial(n-1)
		return result
	}
	return 1
}