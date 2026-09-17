func topKFrequent(nums []int, k int) []int {
	hashTable := make(map[int]int)
	for _, num := range nums {
		hashTable[num]++
	}

	helper := make([][]int, len(nums)+1)
	for key, val := range hashTable {
		helper[val] = append(helper[val], key)
	}

	res := []int{}
	for i := len(helper)-1; i > 0; i-- {
		if helper[i] == nil {
			continue
		}
		for j := range len(helper[i]) {
			res = append(res, helper[i][j])
			if len(res) == k {
				return res
			}
		}
	}
	return res
}