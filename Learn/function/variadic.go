package main

import (
	"fmt"
	"reflect"
)

func variadicExample(s ...string) {
	fmt.Println(s[0], s[2])
}

func variadicExample2(i ...interface{}) {
	for _, val := range i {
		fmt.Println(val, "--", reflect.ValueOf(val).Kind())
	}
}

func calculation(s string, y ...int) int {
	area := 1
	for _, val := range y {
		if s == "Rectangle" {
			area *= val
		} else if s == "Square" {
			area = val * val
		}
	}
	return area
}

func main() {
	variadicExample("martin", "tricia", "cat", "family")
	variadicExample2(1, "Red", []string{"China", "Japan"}, map[string]string{"China": "Beijing", "Japan": "Tokyo"})

	fmt.Println(calculation("Rectangle", 20, 60))
	fmt.Println(calculation("Square", 20, 30, 60))
}
