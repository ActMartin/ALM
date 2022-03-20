package main

import "fmt"

type person struct{
	name string
	age int
}

func student(name string) *person{
	p := person{name:name,age:42}
	return &p
}

func main(){
	fmt.Println(student("martin"))
}