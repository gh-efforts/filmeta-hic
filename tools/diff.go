package tools

func DiffStr(a, b []string) (inAAndB, inAButNotB, inBButNotA []string) {
	m := make(map[string]uint8)
	for _, k := range a {
		m[k] |= 1 << 0
	}
	for _, k := range b {
		m[k] |= 1 << 1
	}

	for k, v := range m {
		x := v&(1<<0) != 0
		y := v&(1<<1) != 0
		switch {
		case x && y:
			inAAndB = append(inAAndB, k)
		case x && !y:
			inAButNotB = append(inAButNotB, k)
		case !x && y:
			inBButNotA = append(inBButNotA, k)
		}
	}
	return
}
