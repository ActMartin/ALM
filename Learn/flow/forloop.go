package main

import "fmt"

func main() {
	k := 1

	for ; k <= 10; k++ {
		fmt.Println(k)
	}
	k = 1
	for k <= 10 {
		fmt.Println(k)
		k++
	}

	k = 1
	for {
		fmt.Println(k)
		k++
		if k > 10 {
			break
		}
	}
}
