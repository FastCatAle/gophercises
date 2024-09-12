package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

type problem struct {
	ques	string
	answ	string
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem {
			ques:	line[0],
			answ:	line[1],
		}
	}
	return ret
}

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
	problems := parseLines(lines)

	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.ques)
		var answer string
		fmt.Scanf("%s/n", &answer)
		if answer == p.answ {
			fmt.Println("Correct!")
			correct++
		}
	}
	fmt.Printf("You scored %d out of %d\n", correct, len(problems))
}
