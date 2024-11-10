package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

/* strategy:
unique domain -> bucket
*/

func main() {
	dataDir := "./data"
	fileNameOrder := []string{"200", "500", "1000", "2000", "5000", "10000", "20000", "50000", "100000", "200000", "500000", "1000000"}
	baseFileName := fmt.Sprintf("%s/top-%s-domains.csv", dataDir, fileNameOrder[0])

	baseFile, err := os.Open(baseFileName)
	reader := csv.NewReader(baseFile)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	// add each record into the set
	domainsMap := make(map[string]string)
	for _, row := range records {
		for _, field := range row {
			domainsMap[field] = "200"
		}
	}

	for _, order := range fileNameOrder {
		// skip the baseFile
		if order == fileNameOrder[0] {
			continue
		}

		// read each file
		currFileName := fmt.Sprintf("%s/top-%s-domains.csv", dataDir, order)
		fmt.Sprintln("Reading file:", currFileName)
		currFile, err := os.Open(currFileName)
		if err != nil {
			log.Fatal(err)
		}
		reader := csv.NewReader(currFile)
		records, err := reader.ReadAll()
		for _, row := range records {
			for _, field := range row {
				// if not already in the map, then add it with the next group
				if _, exists := domainsMap[field]; !exists {
					domainsMap[field] = order
				}

			}
		}
	}

	for k, v := range domainsMap {
		fmt.Printf("%-*s: %s\n", 60, k, v)
	}
}
