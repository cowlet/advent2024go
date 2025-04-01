package safety

import (
	"log"
	"slices"
	"strconv"
	"strings"
)

type Rule struct {
	a int
	b int
}

func ParseRules(lines []string) []Rule {
	var rules []Rule
	for _, s := range lines {
		pgs := strings.Split(s, "|")
		if len(pgs) != 2 {
			log.Fatalf("Rule not <before>|<after>: %s", s)
		}
		a, err := strconv.Atoi(pgs[0])
		if err != nil {
			log.Fatalf("fmt.Atoi: %v", err)
		}
		b, err := strconv.Atoi(pgs[1])
		if err != nil {
			log.Fatalf("fmt.Atoi: %v", err)
		}
		rules = append(rules, Rule{a, b})
	}
	return rules
}

func Validate(book string, rules []Rule) ([]int, bool) {
	pgs := strings.Split(book, ",")
	nos := make([]int, len(pgs))
	for i, p := range pgs {
		n, err := strconv.Atoi(p)
		if err != nil {
			log.Fatalf("page strconv.Atoi: %v", err)
		}
		nos[i] = n
	}

	for _, r := range rules {
		if !passes(r, nos) {
			return nil, false
		}
	}
	return nos, true
}

func passes(r Rule, pgs []int) bool {
	x := slices.Index(pgs, r.a)
	y := slices.Index(pgs, r.b)

	// Are the rule pages even in this set?
	if x == -1 || y == -1 {
		return true // automatic pass
	}

	// Does a occur before b?
	if x < y {
		return true // actual pass
	}
	return false // a must come after b
}
