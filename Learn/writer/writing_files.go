package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	//"strings"
)

func main() {
	// read list of filenames
	f, err := os.Open("info.txt")
	if err != nil {
		log.Fatalln("Couldn't open the filename file", err)
	}
	defer f.Close()

	m := csv.NewReader(f)
	list, err := m.ReadAll()
	if err != nil {
		log.Println("Cannot read CSV file:", err)
	}

	// Create new files
	newFile, err := os.Create("Combined.txt")
	if err != nil {
		log.Fatalln("Couldn't open the filename file", err)
	}
	writer := csv.NewWriter(newFile)
	defer newFile.Close()

	var printed int = 0
	for _, row := range(list) {
		file, err := os.Open(row[0])
		if err != nil {
			log.Fatalln("Couldn't open the csv file", err)
		}

		r := csv.NewReader(file)
		header, _ := r.Read()
		newHeader := append( header,"FILE_NAME",)

		if printed == 0 {
			err = writer.Write(newHeader)
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

			newRecord := append(record, row[1])
			err = writer.Write(newRecord)

		}
	}

	writer.Flush()
}