package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"cowlet.org/advent2024go/day2/reactor"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: day2 <input.txt>")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("os.Open: %v", err)
	}

	// Read lines of space-separated integers. Each line is a Report
	var reports []reactor.Report
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		var r reactor.Report
		for v := range strings.SplitSeq(line, " ") {
			x, err := strconv.Atoi(v)
			if err != nil {
				log.Fatalf("Atoi(%v): %v", v, err)
			}
			r.Append(x)
		}
		//log.Printf("%v", r)
		reports = append(reports, r)
	}

	count := 0
	for _, r := range reports {
		if r.Safe() {
			count += 1
		}
		//log.Printf("report %v safe? %v", r, r.Safe())
	}
	log.Printf("Total number of safe reports: %d", count)
}
