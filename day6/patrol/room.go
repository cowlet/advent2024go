package room

import (
	"log"
	"slices"
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

	log.Print("Running guard...")
	r.runGuard()
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

func (r *Room) runGuard() {
	var i, j int
	var c rune
	for {
		switch curr := r.g.path[0]; curr.c {
		case '^':
			i = curr.x - 1
			j = curr.y
			c = '^'
			if r.CheckObs(i, j) != nil {
				i = curr.x
				c = '>'
			}
		case 'v':
			i = curr.x + 1
			j = curr.y
			c = 'v'
			if r.CheckObs(i, j) != nil {
				i = curr.x
				c = '<'
			}
		case '>':
			i = curr.x
			j = curr.y + 1
			c = '>'
			if r.CheckObs(i, j) != nil {
				j = curr.y
				c = 'v'
			}
		case '<':
			i = curr.x
			j = curr.y - 1
			c = '<'
			if r.CheckObs(i, j) != nil {
				j = curr.y
				c = '^'
			}
		}
		if i < 0 || i >= r.nr || j < 0 || j >= r.nc {
			log.Printf("Guard exited at (%d, %d)\n", i, j)
			return
		}
		next := Coord{i, j, c}
		r.g.path = append([]Coord{next}, r.g.path...)
	}
}

func (r *Room) CountGuard() int {
	// Ignore the rune and count only unique positions
	var uniques []Coord
	for _, step := range r.g.path {
		exists := slices.ContainsFunc(uniques, func(a Coord) bool {
			return a.x == step.x && a.y == step.y
		})
		if !exists {
			uniques = append(uniques, step)
		}

	}
	return len(uniques)
}
