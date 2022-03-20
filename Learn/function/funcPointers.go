package main

import "fmt"

func update(i *int, j *string) {
	*i = *i + 5
	*j = *j + "Doe"
	return
}

func main() {
	age := 20
	name := "John"
	fmt.Println("Before: ", age, name)
	update(&age, &name)
	fmt.Println("After: ", age, name)
}
