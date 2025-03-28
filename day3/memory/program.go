package memory

import (
	"cmp"
	"log"
	"regexp"
	"slices"
	"strconv"
)

type Program struct {
	muls  []Mul
	dos   []Do
	donts []Dont
}

type Instruction interface {
	Pos() int
}

type Mul struct {
	pos int
	x   int
	y   int
}

func (m Mul) Pos() int {
	return m.pos
}

func parsemul(s string, loc []int) *Mul {
	x, err := strconv.Atoi(s[loc[2]:loc[3]])
	if err != nil {
		log.Fatalf("strconv.Atoi: %v", err)
	}
	y, err := strconv.Atoi(s[loc[4]:loc[5]])
	if err != nil {
		log.Fatalf("strconv.Atoi: %v", err)
	}
	m := Mul{loc[0], x, y}
	return &m
}

type Do struct {
	pos int
}

func (d Do) Pos() int {
	return d.pos
}

func parsedo(loc []int) *Do {
	d := Do{loc[0]}
	return &d
}

type Dont struct {
	pos int
}

func (d Dont) Pos() int {
	return d.pos
}

func parsedont(loc []int) *Dont {
	d := Dont{loc[0]}
	return &d
}

func (prog *Program) Parse(s string) {
	mul := regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)`)
	do := regexp.MustCompile(`do\(\)`)
	dont := regexp.MustCompile(`don't\(\)`)

	allmuls := mul.FindAllStringSubmatchIndex(s, -1)
	for _, match := range allmuls {
		prog.muls = append(prog.muls, *parsemul(s, match))
	}

	alldos := do.FindAllStringIndex(s, -1)
	for _, match := range alldos {
		prog.dos = append(prog.dos, *parsedo(match))
	}

	alldonts := dont.FindAllStringIndex(s, -1)
	for _, match := range alldonts {
		prog.donts = append(prog.donts, *parsedont(match))
	}
}

func (p *Program) next() *Instruction {
	var compare []Instruction
	if len(p.muls) > 0 {
		compare = append(compare, p.muls[0])
	}
	if len(p.dos) > 0 {
		compare = append(compare, p.dos[0])
	}
	if len(p.donts) > 0 {
		compare = append(compare, p.donts[0])
	}

	next := slices.MinFunc(compare, func(a, b Instruction) int {
		return cmp.Compare(a.Pos(), b.Pos())
	})
	return &next
}

func (p *Program) Execute() int {
	// Need to find the correct execution order
	total := 0
	on := true

	for {
		if len(p.muls) == 0 {
			return total
		}
		next := *p.next()

		// For each type of instruction, execute it then consume it
		switch next.(type) {
		case Mul:
			if on {
				total += p.muls[0].x * p.muls[0].y
			}
			p.muls = p.muls[1:]
		case Do:
			on = true
			p.dos = p.dos[1:]
		case Dont:
			on = false
			p.donts = p.donts[1:]
		}
	}
}
