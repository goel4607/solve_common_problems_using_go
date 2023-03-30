package Q4_FindTypedOutStrings

type TypedOutStringsEqualsFinderS1Bf struct {
}

func (bf TypedOutStringsEqualsFinderS1Bf) findIfTypedOutStringsAreEqual(s, t string) bool {
	s1 := bf.findTypedOutString(s)
	t1 := bf.findTypedOutString(t)

	if len(s1) != len(t1) {
		return false
	}

	for i := range s1 {
		if s1[i] != t1[i] {
			return false
		}
	}

	return true
}

func (bf TypedOutStringsEqualsFinderS1Bf) findTypedOutString(str string) string {
	s := stack{chars: make([]rune, 0)}
	for i := range str {
		if str[i] == hash {
			s.pop() // ignore the return value
		} else {
			s.push(rune(str[i]))
		}
	}

	return string(s.chars)
}

type stack struct {
	chars []rune
}

func (s *stack) push(v rune) {
	s.chars = append(s.chars, v)
}

func (s *stack) pop() rune {
	if len(s.chars) > 0 {
		c := s.chars[len(s.chars)-1]
		s.chars = s.chars[0 : len(s.chars)-1]
		return c
	}

	return -1
}
