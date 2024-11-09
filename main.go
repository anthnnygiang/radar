package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

/* Strategy:
Read top 1,000,000 domains
Add them into a set
Read top 500,000 domains
If domain is in the 1,000,000 set
then remove it and add it to the 500,000 set
repeat.
*/

func processFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
}

func main() {
	files, err := os.ReadDir("./data")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	set := make(map[string]struct{})

	for _, row := range records {
		for _, field := range row {
			fmt.Println(field)
			set[field] = struct{}{}
		}
	}

}
