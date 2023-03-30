package Q24_Course_scheduler

type CourseDependency struct {
	course int
	preReq int
}

type Interface interface {
	isItPossibleToFinishAllCourses(n int, prereqs []CourseDependency) bool
}
