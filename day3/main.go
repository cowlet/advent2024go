package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	"cowlet.org/advent2024go/day3/memory"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: day3 <input.txt>")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("os.Open: %v", err)
	}

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		//log.Printf("line is: %s", line)
	}

	text := strings.Join(lines, "")
	var prog memory.Program
	prog.Parse(text)
	//log.Printf("%v", prog)
	log.Printf("Program returns %d", prog.Execute())
}
