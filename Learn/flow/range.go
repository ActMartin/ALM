package main

import "fmt"

func main() {
	varDict := map[string]string{"China": "Beijing", "Japan": "Tokyo", "Canada": "Ottawa"}

	for idx, element := range varDict {
		fmt.Println("Index : ", idx, "Element : ", element)
	}

	for range "HELLO" {
		fmt.Println("Hello")
	}
}
