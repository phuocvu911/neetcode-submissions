func searchMatrix(matrix [][]int, target int) bool {
	for _, slice := range matrix{
		for _, num := range slice{
			if num == target{
				return true
			}
		}
	}
	return false
}
