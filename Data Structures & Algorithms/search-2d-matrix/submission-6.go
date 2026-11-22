func searchMatrix(matrix [][]int, target int) bool {
	m, n:= len(matrix), len(matrix[0])
	head, tail:= 0, m*n -1
	for head <= tail{
		mid := head + (tail-head)/2
		i,j := mid/n, mid%n

		if matrix[i][j] > target{
			tail = mid -1
		} else if matrix[i][j] < target{
			head = mid + 1
		} else{
			return true
		}
	}
	return false
}
