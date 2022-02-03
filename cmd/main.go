package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	absPath, err := filepath.Abs("Higload-L20-Backups/books.csv")
	if err != nil {
		log.Fatalf("Error: '%v'\n", err)
	}
	records := readCsvFile(absPath)
	fmt.Println(records)
}

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}
