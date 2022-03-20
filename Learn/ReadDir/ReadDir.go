package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	myfile, err := os.Create("myfile.txt")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fi,_:=os.Stat(file.Name())
		martin := io.MultiWriter(os.Stdout, myfile)
		fmt.Fprintln(martin, file.Name(),",", fi.Size())
	}
}
