package Q24_Course_scheduler

type Soln2TopSort struct {
}

func (s Soln2TopSort) canFinish(n int, preReqs [][2]int) bool {
	if n <= 1 {
		return true
	}
	courseInDegMap := s.convertToInDegreeMaps(n, preReqs)

	visitedZeroDegCourses := make(map[int]bool, 0)

	for {
		beforeZeroDeg := len(visitedZeroDegCourses)
		for i := 0; i < len(courseInDegMap); i++ {
			if visitedZeroDegCourses[i] {
				continue
			}
			m := courseInDegMap[i]
			if len(m) == 0 {
				for j := 0; j < len(courseInDegMap); j++ {
					if i == j || visitedZeroDegCourses[i] {
						continue
					}

					m1 := courseInDegMap[j]
					if len(m1) > 0 && m1[i] {
						delete(m1, i)
					}
				}

				visitedZeroDegCourses[i] = true
			}
		}

		if len(visitedZeroDegCourses) == n {
			break
		}

		if beforeZeroDeg == len(visitedZeroDegCourses) { //no change means no new zero deg found so there must be a cycle in it
			return false
		}
	}

	return true
}

func (s Soln2TopSort) convertToInDegreeMaps(n int, preReqs [][2]int) []map[int]bool {
	courseInDegMap := make([]map[int]bool, n, n)

	for _, p := range preReqs {
		inDegMap := courseInDegMap[p[0]]
		if courseInDegMap[p[0]] == nil {
			inDegMap = make(map[int]bool)
			courseInDegMap[p[0]] = inDegMap
		}

		inDegMap[p[1]] = true
	}

	return courseInDegMap
}
