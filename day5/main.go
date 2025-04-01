package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	"cowlet.org/advent2024go/day5/safety"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: day5 <input.txt>")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("os.Open: %v", err)
	}

	var rlines []string
	var blines []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.ContainsRune(line, '|') {
			rlines = append(rlines, line)
		} else if strings.ContainsRune(line, ',') {
			blines = append(blines, line)
		}
	}

	rules := safety.ParseRules(rlines)

	valids := make([][]int, 0, len(blines))
	invalids := make([][]int, 0, len(blines))
	for _, b := range blines {
		pgs, valid := safety.Validate(b, rules)
		if valid {
			valids = append(valids, pgs)
		} else {
			invalids = append(invalids, pgs)
		}
	}
	log.Printf("Total number of valid books: %d", len(valids))

	total := 0
	for _, v := range valids {
		total += v[len(v)/2]
	}
	log.Printf("Sum of the valid middle page numbers: %d", total)

	total = 0
	for _, pgs := range invalids {
		safety.Fix(pgs, rules)
		total += pgs[len(pgs)/2]
	}
	log.Printf("Sum of the corrected invalid middle page numbers: %d", total)
}
