package main

import (
	"bufio"
	"log"
	"os"

	room "cowlet.org/advent2024go/day6/patrol"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: day6 <input.txt>")
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

	r := room.NewRoom(lines)
	r.Print()
	log.Printf("The guard visited %d locations", r.CountGuard())
}
