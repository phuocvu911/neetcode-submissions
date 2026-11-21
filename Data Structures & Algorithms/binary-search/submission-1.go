func search(nums []int, target int) int {
	head, tail := 0, len(nums)-1
	for head <= tail{
		checker := head + (tail-head)/2
		if nums[checker] < target{
			head = checker + 1
		} else if nums[checker] > target{
			tail = checker - 1
		} else {
			return checker
		}
	}
	return -1
}
