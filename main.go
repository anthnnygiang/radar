package main

import (
	"encoding/csv"
	"fmt"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"log"
	"os"
	"path/filepath"
)

func main() {
	dataDir := "data"
	fileFormat := "cloudflare-radar_top-%d-domains_.csv"
	brackets := []int{200, 500, 1000, 2000, 5000, 10000, 20000, 50000, 100000, 200000, 500000, 1000000}
	domainsMap := make(map[string]int)
	p := message.NewPrinter(language.English)
	file, err := os.Create("result.txt")
	if err != nil {
		log.Fatal(err)
	}

	// read each file
	for _, bracket := range brackets {
		currFileName := filepath.Join(dataDir, fmt.Sprintf(fileFormat, bracket))
		fmt.Println("Reading file:", currFileName)
		currFile, err := os.Open(currFileName)
		if err != nil {
			log.Fatal(err)
		}
		reader := csv.NewReader(currFile)
		// read the first row to ignore headers
		_, err = reader.Read()
		if err != nil {
			log.Fatal(err)
		}

		records, err := reader.ReadAll()
		for _, row := range records {
			for _, field := range row {
				// if not already in the map, then add it and write to file
				if _, exists := domainsMap[field]; !exists {
					domainsMap[field] = bracket
					_, err := file.Write([]byte(p.Sprintf("%*d: %s\n", 7, bracket, field)))
					if err != nil {
						log.Fatal(err)
					}
				}

			}
		}
	}
}
