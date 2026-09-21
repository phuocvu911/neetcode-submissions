func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	numSet := make(map[int]bool, len(nums))
	for _, num := range nums {
		numSet[num] = true
	}
	maxLength := 0
	for num := range numSet {
		if !numSet[num-1] {
			currentNum := num
			length := 1
			for numSet[currentNum+1] {
				currentNum++
				length++

			}
			if length > maxLength {
				maxLength = length
			}
		}
	}
	return maxLength
}