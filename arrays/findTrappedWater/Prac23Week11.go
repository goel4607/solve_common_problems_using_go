package findTrappedWater

//start time: 7:10 pm
//start time: 7:42 pm

type Prac23Week11 struct {
}

func (p Prac23Week11) findTrappedWater(w []int) int {
	if len(w) <= 2 {
		return 0
	}

	var total int

	//use two pointers say left and right and init them with 0 and len() -1
	left, right := 0, len(w)-1
	//declare total, maxLeft and maxRight
	maxLeft, maxRight := w[left], w[right]
	//iterate wile left < right
	left++
	right--
	for left <= right {
		//check which element is <= out of left and right
		if maxLeft <= maxRight {
			//watar can only be trapped if left wall is > 0 and > this element
			if maxLeft > 0 && maxLeft > w[left] {
				diff := maxLeft - w[left]
				total += diff
			}

			//update max if it is < than current value
			if maxLeft < w[left] {
				maxLeft = w[left]
			}
			left++
		} else { //do the same for right
			if maxRight > 0 && maxRight > w[right] {
				diff := maxRight - w[right]
				total += diff
			}

			if maxRight < w[right] {
				maxRight = w[right]
			}

			right--
		}
	}

	//return total
	return total
}
