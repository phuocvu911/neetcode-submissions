func productExceptSelf(nums []int) []int {
	product := 1
	countZero := 0
	for _, num := range nums {
		if num == 0 {
			countZero++
			continue
		}
		product *= num
	}
	result := make([]int, len(nums))

	if countZero > 1 {
		return result
	}
	if countZero == 1 {
		for i, num := range nums {
			if num == 0 {
				result[i] = product
				continue
			}
			result[i] = 0
		}
		return result
	}
	for i, num := range nums {
		result[i] = product / num
	}
	return result
}
