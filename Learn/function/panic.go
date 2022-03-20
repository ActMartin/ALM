package main

import "fmt"

func main() {
	var action int
	fmt.Println("Enter 1 for Student and 2 for Professional")
	fmt.Scanln(&action)

	switch action {
	case 1:
		fmt.Println("I am a student")
	case 2:
		fmt.Println("I am a professional")
	default:
		panic(fmt.Sprintf("I am a %d", action))
	}
}
