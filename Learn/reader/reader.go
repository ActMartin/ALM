package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type Epl struct {
	idxCflow int
	varName string
	varValue string
}

func main() {
	// read map
	mapVar := make(map[string] int)
	mapfile, err := os.Open("EPL_VARS.txt")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	m := csv.NewReader(mapfile)

	records, err := m.ReadAll()
	if err != nil {
		log.Println("Cannot read CSV file:", err)
	}
	for _, row := range(records) {
		id, _ := strconv.ParseInt(row[0], 0, 0)
		mapVar[row[1]] = int(id)
	}

	// Open the file
	csvfile, err := os.Open("EPL_PROJ_201912_BASE.fac")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)
	var epls []Epl
	var arr [][]int

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

		idx, _ := strconv.ParseInt(record[2],0,0)
		idx2, _:= strconv.ParseInt(record[4],0,0)
		epl := Epl{int(idx),record[3],record[4]}

		epls = append(epls,epl)

		arr[int(idx)][mapVar[record[3]]] = int(idx2)

	}

	fmt.Println(arr[1][mapVar["NO_POLS_IF"]])
	fmt.Println(1, epls[1].idxCflow, epls[1].varName)

}