func searchRange(nums []int, target int) []int {
	l := sort.SearchInts(nums, target)
	r := sort.SearchInts(nums, target+1)
	if l == r {
		return []int{-1, -1}
	}
	return []int{l, r - 1}
}

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	left := searchLeftIndex(nums, target)
	if left == nums[len(nums)-1] || nums[left] != target {
		return []int{-1, -1}
	}
	right := searchLeftIndex(nums, target)
	return []int{left, right}
}

func searchLeftIndex(nums []int, target int) int {
	l, r := 0, len(nums)

	for l < r {
		mid := l + (r-l)/2
		if nums[l] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

func searchRightIndex(nums []int, target int) int {
	l, r := 0, len(nums)

	for l < r {
		mid := l + (r-l)/2
		if nums[l] <= target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}
