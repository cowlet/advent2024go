package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: day3 <input.txt>")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("os.Open: %v", err)
	}

	var mulexp = regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)`)
	total := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		//log.Printf("line is: %s", line)
		matches := mulexp.FindAllStringSubmatch(line, -1)
		for _, s := range matches {
			//log.Printf("First match is %v", s)
			x, err := strconv.Atoi(s[1])
			if err != nil {
				log.Fatalf("strconv.Atoi: %v", err)
			}
			y, err := strconv.Atoi(s[2])
			if err != nil {
				log.Fatalf("strconv.Atoi: %v", err)
			}
			//log.Printf("Found %d * %d", x, y)
			total += x * y
		}
	}
	log.Printf("Final total is %d", total)
}
