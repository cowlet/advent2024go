package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: day1 <file.txt>")
	}
	f, err := os.Open(os.Args[1]) // cmd line arg
	if err != nil {
		log.Fatalf("os.Open: %v", err)
	}

	// Read position values into two sorted lists
	var xs []int
	var ys []int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		var x int
		var y int
		_, err := fmt.Sscanf(line, "%d   %d", &x, &y)
		if err != nil {
			log.Fatalf("scanning %q: %v", line, err)
		}
		xs = insert(xs, x) // insertion sort
		ys = insert(ys, y)
	}

	// Now calculate distances between pairs of positions
	distance := 0
	for i := range xs {
		distance = distance + absdistance(xs[i], ys[i])
	}
	log.Printf("The total distance is %d", distance)
}

func insert(s []int, val int) []int {
	for i, v := range s {
		if v >= val {
			// Insert at i and push everything else up
			return slices.Insert(s, i, val)
		}
	}
	// If we get here, pos is the largest, so add to the end
	return append(s, val)
}

func absdistance(x int, y int) int {
	if d := x - y; d < 0 {
		return -d
	} else {
		return d
	}
}
