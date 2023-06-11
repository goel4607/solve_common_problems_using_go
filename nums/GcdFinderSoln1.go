package nums

type GcdFinderSoln1 struct {
}

func (s GcdFinderSoln1) findGcd(a, b int) int {

	for {
		r := a % b
		a, b = b, r
		if r == 0 {
			break
		}
	}

	return a
}
