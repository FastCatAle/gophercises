package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFilename := flag.String("file", "problems.csv", "CSV file to read in the format 'question,answer'")
	flag.Parse()
	file, err := os.Open(*csvFilename)
	if err != nil {
		fmt.Printf("Failed to open CSV file: %s\n", *csvFilename)
		os.Exit(1)
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		fmt.Println("Failed to read the CSV file.")
		os.Exit(1)
	}
	fmt.Println(lines)
}
