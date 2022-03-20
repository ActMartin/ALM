package main

import (
	"encoding/csv"
	"strconv"

	//"strconv"
	"fmt"
	//"io"
	"log"
	"os"
)


type User struct {
	id int
	firstName string
	lastName string
	email string
}

func main(){
	file,err := os.Open("users.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		log.Println("Cannot read CSV file:", err)
	}

	for _, row := range(records){
		id, _ := strconv.ParseInt(row[0],0,0)
		user := User{int(id),row[1],row[2],row[3]}

		fmt.Println(user)
	}
}