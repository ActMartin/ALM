package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

func main() {
	// read list of filenames
	f, err := os.Open("filenames.txt")
	if err != nil {
		log.Fatalln("Couldn't open the filename file", err)
	}
	defer f.Close()

	newFile, err := os.Create("Combined.txt")
	if err != nil {
		log.Fatalln("Couldn't open the filename file", err)
	}
	writer := csv.NewWriter(newFile)
	defer newFile.Close()

	m := csv.NewReader(f)

	list, err := m.ReadAll()
	if err != nil {
		log.Println("Cannot read CSV file:", err)
	}

	var printed int = 0
	for _, row := range(list) {
		file, err := os.Open(row[0])
		if err != nil {
			log.Fatalln("Couldn't open the csv file", err)
		}

		r := csv.NewReader(file)
		header, _ := r.Read()
		if printed == 0 {
			err = writer.Write(header)
			printed = 1
		}

		// Iterate through the records
		for {
			// Read each record from csv
			record, err := r.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}

			err = writer.Write(record)

		}
	}

	writer.Flush()
}