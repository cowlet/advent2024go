package main

import (
	"bufio"
	"log"
	"os"

	"cowlet.org/advent2024go/day4/grid"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: day4 <input.txt>")
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("os.Open: %v", err)
	}
	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	s := grid.NewSearch("XMAS", lines)
	log.Printf("Total number of XMASes is %d", s.CountNormal())
	log.Printf("Total number of X-MASes is %d", s.CountXMas())
}
