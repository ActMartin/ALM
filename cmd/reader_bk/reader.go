package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	//Slice for storing EPL data
	var Epl [11][34][1201]string

	// Map Var
	mapVar := make(map[string]int)
	mapFile, err := os.Open("EPL_VARS.txt")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	m := csv.NewReader(mapFile)

	records, err := m.ReadAll()
	if err != nil {
		log.Println("Cannot read CSV file:", err)
	}
	for _, row := range records {
		id, _ := strconv.ParseInt(row[0], 0, 0)
		mapVar[row[1]] = int(id)
	}
	fmt.Println(mapVar)

	//Map idx
	mapCfs := make(map[string]int)
	mapFile, err = os.Open("EPL_CFS.txt")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	m = csv.NewReader(mapFile)

	records, err = m.ReadAll()
	if err != nil {
		log.Println("Cannot read CSV file:", err)
	}
	for _, row := range records {
		id, _ := strconv.ParseInt(row[0], 0, 0)
		mapCfs[row[1]] = int(id)
	}
	fmt.Println(mapCfs)

	//Map Caldate
	mapYYYYMM := make(map[string]int)
	mapFile, err = os.Open("MAP_YYYYMM.txt")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	m = csv.NewReader(mapFile)

	records, err = m.ReadAll()
	if err != nil {
		log.Println("Cannot read CSV file:", err)
	}
	for _, row := range records {
		id, _ := strconv.ParseInt(row[0], 0, 0)
		mapYYYYMM[row[1]] = int(id)
	}
	fmt.Println(mapYYYYMM)

	csvFile, err := os.Open("EPL_PROJ_201912_Base.fac")
	if err != nil {
		log.Fatalln(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)

	var idxCflow string

	idxExistMap := make(map[string]bool)
	//datExistMap := make(map[string]bool)

	var IntIdx int
	for {
		// Iterate through records
		record, err := reader.Read()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}

		idxCflow = record[2] + record[3]
		if _, exist := idxExistMap[idxCflow]; exist {
			log.Fatalln("duplicate index found for ", idxCflow, err)
		} else {
			idxExistMap[idxCflow] = true
		}

		//
		start := record[0]
		if strings.HasPrefix(record[0], "!") {
			//Parse header
			//fmt.Println(record)
			fmt.Println(start)
			Num := strings.TrimPrefix(record[0], "!")
			Idx, _ := strconv.ParseInt(Num, 0, 0)
			IntIdx = int(Idx)
		} else {
			i := mapCfs[record[2]]
			j := mapVar[record[3]]
			for k, value := range record[IntIdx:1200] {
				//k := mapYYYYMM[yyyymm]
				Epl[i][j][k] = value
				//fmt.Println(k, record[2], record[3], value, Epl[i][j][k])
			}
		}
		fmt.Println(IntIdx)
	}
	fmt.Println(Epl[1][1][2], "end")
}
