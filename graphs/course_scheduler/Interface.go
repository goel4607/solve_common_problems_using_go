package course_scheduler

type CourseDependency struct {
	course int
	preReq int
}

type Interface interface {
	isItPossibleToFinishAllCourses(n int, prereqs []CourseDependency) bool
}
