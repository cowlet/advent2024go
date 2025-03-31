package grid

import (
	"log"
	"strings"
	"unicode/utf8"
)

type Search struct {
	Text string
	Rows [][]rune
	nr   int
	nc   int
}

func NewSearch(text string, lines []string) *Search {
	if len(lines) < 1 {
		log.Fatal("Cannot search with no lines")
	}
	nr := len(lines)
	nc := utf8.RuneCountInString(lines[0])
	for i, row := range lines {
		c := utf8.RuneCountInString(row)
		if c != nc {
			log.Fatalf("Row %d has %d runes instead of %d", i, c, nc)
		}
	}
	s := Search{Text: text, nr: nr, nc: nc}
	s.Rows = make([][]rune, nr)
	for i, line := range lines {
		s.Rows[i] = []rune(line)
	}
	return &s
}

func (s *Search) Count() int {
	a := s.countRows()
	b := s.countCols()
	c := s.countLDiag()
	d := s.countRDiag()
	return a + b + c + d
}

func (s *Search) countRows() int {
	total := 0
	for _, fwd := range s.Rows {
		total += strings.Count(string(fwd), s.Text)

		// Reverse, and count again
		bck := make([]rune, s.nc)
		for i, ch := range fwd {
			bck[s.nc-1-i] = ch
		}
		total += strings.Count(string(bck), s.Text)
	}
	return total
}

func (s *Search) countCols() int {
	total := 0
	for j := range s.nc {
		dwn := make([]rune, s.nr)
		for i := range s.nr {
			dwn[i] = s.Rows[i][j]
		}
		total += strings.Count(string(dwn), s.Text)

		// Reverse, and count again
		up := make([]rune, s.nr)
		for i, ch := range dwn {
			up[s.nr-1-i] = ch
		}
		total += strings.Count(string(up), s.Text)
	}
	return total
}

func (s *Search) countLDiag() int {
	total := 0
	// Lower left diagonal
	for i := range s.nr {
		diag := make([]rune, max(s.nc, s.nr))
		for j := range s.nc {
			if i+j >= s.nr {
				break // overrunning grid edge
			}
			diag[j] = s.Rows[i+j][j]
		}
		total += strings.Count(string(diag), s.Text)

		// Reverse, and count again
		n := len(diag)
		rev := make([]rune, n)
		for k, ch := range diag {
			rev[n-1-k] = ch
		}
		total += strings.Count(string(rev), s.Text)
	}

	// Upper right diagonal. No need to recount (0, 0), so start at j=1
	for j := 1; j < s.nc; j++ {
		diag := make([]rune, max(s.nr, s.nc))
		for i := range s.nr {
			if i+j >= s.nc {
				break // overrunning grid edge
			}
			diag[i] = s.Rows[i][i+j]
		}
		total += strings.Count(string(diag), s.Text)

		// Reverse, and count again
		n := len(diag)
		rev := make([]rune, n)
		for k, ch := range diag {
			rev[n-1-k] = ch
		}
		total += strings.Count(string(rev), s.Text)
	}
	return total
}

func (s *Search) countRDiag() int {
	total := 0
	// Lower right diagonal
	for i := range s.nr {
		diag := make([]rune, max(s.nr, s.nc))
		for j := range s.nc {
			if i+j >= s.nr {
				break
			}
			diag[j] = s.Rows[i+j][s.nc-1-j]
		}
		total += strings.Count(string(diag), s.Text)

		// Reverse, and count again
		n := len(diag)
		rev := make([]rune, n)
		for k, ch := range diag {
			rev[n-1-k] = ch
		}
		total += strings.Count(string(rev), s.Text)
	}

	// Upper left diagonal. No need to recount (0, s.nc-1), so end at j > 0
	for j := s.nc - 1; j > 0; j-- {
		diag := make([]rune, max(s.nr, s.nc))
		for i := range s.nr {
			if i+j >= s.nc {
				break
			}
			diag[i] = s.Rows[i][s.nc-1-j-i]
		}
		total += strings.Count(string(diag), s.Text)

		// Reverse, and count again
		n := len(diag)
		rev := make([]rune, n)
		for k, ch := range diag {
			rev[n-1-k] = ch
		}
		total += strings.Count(string(rev), s.Text)
	}
	return total
}
