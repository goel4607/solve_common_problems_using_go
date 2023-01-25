package cycledetector

type CycleDetector interface {
	detectCycle(head *CNode) bool
}
