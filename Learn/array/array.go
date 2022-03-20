package main

import "fmt"

func main(){
	var array1 [10] int
	var i int
	for i = 0; i < 10; i++{
		array1[i] = i + 100
	}
	//output
	for i = 0; i < 10; i++{
		fmt.Printf("Element[%d] = %d\n", i, array1[i])
	}


}