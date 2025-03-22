package reactor

import "slices"

type Report struct {
	a []int
}

func (r *Report) Append(v int) {
	r.a = append(r.a, v)
}

func (r *Report) Safe() bool {
	return r.fullsafe() || r.dampener()
}

func (r *Report) fullsafe() bool {
	// Condition 1: is monotonic increasing
	monoinc := slices.IsSorted(r.a)
	if monoinc {
		// Condition 2: Jumps are positive steps of 1, 2, or 3
		return safesteps(r.a)
	}

	// Might still be monotonically decreasing
	monodec := slices.IsSortedFunc(r.a, func(x, y int) int {
		if x == y {
			return 0
		} else if x < y {
			return 1
		} else {
			return -1
		}
	})
	// Neater code, but double allocation if not monoinc
	// rev := slices.Clone(r.a)
	// slices.Reverse(rev)
	// monodec := slices.IsSorted(rev)
	if monodec {
		return safesteps(r.a)
	}

	return false // Not monotonic
}

func safesteps(ss []int) bool {
	for i := 1; i < len(ss); i++ {
		step := ss[i] - ss[i-1]
		if step < 0 {
			step = step * -1
		}
		if step < 1 || step > 3 {
			return false
		}
	}
	return true
}

// dampener checks whether removing any one entry from the report makes it safe
func (r *Report) dampener() bool {
	for i := range r.a {
		d := Report{slices.Concat(r.a[:i], r.a[i+1:])}
		if d.fullsafe() {
			return true
		}
	}
	return false
}
