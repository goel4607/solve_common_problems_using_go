package Q24_Course_scheduler

//type CourseDependency struct {
//	course int
//	preReq int
//}

type Interface interface {
	canFinish(numCourses int, prerequisites [][]int) bool
}
