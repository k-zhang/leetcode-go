package easy

type RangeSumQuery struct {
	accSum []int
}

func NewRangeSumQuery(nums []int) *RangeSumQuery {
	r := RangeSumQuery{make([]int, len(nums))}
	r.accSum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		r.accSum[i] = r.accSum[i-1] + nums[i]
	}
	return &r
}

func (r *RangeSumQuery) SumRange(left int, right int) int {
	var leftValue int
	if left == 0 {
		leftValue = 0
	} else {
		leftValue = r.accSum[left-1]
	}

	return r.accSum[right] - leftValue
}
