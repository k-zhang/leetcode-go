package easy

import "errors"

func twoSum(nums []int, target int) (int, int, error) {
	dict := map[int]int{}

	for i := 0; i < len(nums); i++ {
		want := target - nums[i]

		_, prs := dict[want]
		if prs {
			return dict[want], i, nil
		}

		dict[nums[i]] = i
	}

	return -1, -1, errors.New("unable to find two numbers that sum up to the target")
}
