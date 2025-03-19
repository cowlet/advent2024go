package main

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Report struct {
	a []int
}

func (r Report) Safe() bool {
	// Is the full thing safe?
	if r.fullsafe() {
		return true
	}
	// If not, try the dampener
	return r.dampener()
}

func (r Report) fullsafe() bool {
	// Condition 1: is monotonic increasing
	monoinc := slices.IsSorted(r.a)
	if monoinc {
		// Condition 2: Jumps are positive steps of 1, 2, or 3
		return safesteps(r.a)
	}

	// Might still be monotonically decreasing
	rev := slices.Clone(r.a)
	slices.Reverse(rev)
	monodec := slices.IsSorted(rev)
	if monodec {
		// Condition 2 on rev
		return safesteps(rev)
	}

	return false // Not monotonic
}

func safesteps(s []int) bool {
	// Can assume positive increments
	for i := 1; i < len(s); i++ {
		step := s[i] - s[i-1]
		if step < 1 || step > 3 {
			return false
		}
	}
	return true
}

func (r Report) dampener() bool {
	// If we remove any one entry from the report, does it become safe?
	for i := range r.a {
		d := Report{slices.Concat(r.a[:i], r.a[i+1:])}
		if d.fullsafe() {
			return true
		}
	}
	return false
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: day2 <input.txt>")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("os.Open: %v", err)
	}

	// Read lines of space-separated integers. Each line is a Report
	var reports []Report
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		r := Report{}
		parts := strings.Split(line, " ")
		r.a = make([]int, len(parts)) // Make it the correct size
		for i, v := range parts {
			r.a[i], err = strconv.Atoi(v)
			if err != nil {
				log.Fatalf("Atoi(%v): %v", v, err)
			}
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
