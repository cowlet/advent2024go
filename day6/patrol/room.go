package room

import (
	"log"
	"unicode/utf8"
)

type Coord struct {
	x int
	y int
	c rune
}

type Guard struct {
	path []Coord
}

type Room struct {
	nr  int
	nc  int
	obs []Coord
	g   Guard
}

func NewRoom(lines []string) *Room {
	if len(lines) < 1 {
		log.Fatalf("need more than %d rows in room", len(lines))
	}
	if utf8.RuneCountInString(lines[0]) < 1 {
		log.Fatalf("need more than %d cols in room", utf8.RuneCountInString(lines[0]))
	}
	r := Room{nr: len(lines), nc: utf8.RuneCountInString(lines[0])}

	for i, line := range lines {
		if utf8.RuneCountInString(line) != r.nc {
			log.Fatalf("wrong length of row: line %d", i)
		}
		for j, ch := range []rune(line) {
			if ch == '#' {
				r.obs = append(r.obs, Coord{i, j, '#'})
			} else if ch == '^' {
				r.g = Guard{[]Coord{{i, j, '^'}}}
			}
		}
	}
	return &r
}

func (r *Room) CheckObs(i int, j int) *rune {
	for _, c := range r.obs {
		if c.x == i && c.y == j {
			return &c.c
		}
	}
	return nil
}

func (r *Room) CheckGuard(i int, j int) *rune {
	for _, c := range r.g.path {
		if c.x == i && c.y == j {
			return &c.c
		}
	}
	return nil
}

func (r *Room) Print() {
	s := make([]rune, r.nr*r.nc+r.nr) // extra for \n
	p := 0
	for i := range r.nr {
		for j := range r.nc {
			if ch := r.CheckObs(i, j); ch != nil {
				s[p] = *ch
			} else if ch := r.CheckGuard(i, j); ch != nil {
				s[p] = *ch
			} else {
				s[p] = '.'
			}
			p++
		}
		s[p] = '\n'
		p++
	}
	log.Printf("\n%s", string(s))
}
