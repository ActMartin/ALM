package main

import (
	"fmt"
	//"github.com/ActMartin/alm/cmd"
	"learn/slice/cmd"
	"math/rand"
	"runtime"
	"time"
)

func makeArr() []int {
	arr := make([]int, 8000000)
	return arr
}
func main() {
	arr := makeArr()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(arr); i++ {
		arr[i] = 100
	}
	arr = arr[:5]
	cmd.PrintMem() //61.218254 MB
	runtime.GC()
	cmd.PrintMem() //0.184059 MB
	fmt.Println(len(arr))
	//arr[0] = 1
}
