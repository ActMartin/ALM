package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	information, _ := os.Create("info.txt")

	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
				return err
			}

			//fmt.Fprintln(io.MultiWriter(os.Stdout, information), path,"||", info.Size(), "||", info.ModTime())

			if !info.IsDir() && strings.Contains(info.Name(), "fac"){
				fmt.Fprintln(io.MultiWriter(os.Stdout, information), path, ",", strings.TrimSuffix(info.Name(),".fac"))
			}

			return nil
		})

	if err != nil {
		log.Println(err)
	}
}

