package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	strVar := "100"
	intVar, _ := strconv.Atoi(strVar)

	fmt.Println(strVar, intVar, reflect.TypeOf(intVar))
}
