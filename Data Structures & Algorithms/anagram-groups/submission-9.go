
func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0{
		return nil
	}

	hashTable := make(map[[26]int][]string)

	for _, str := range strs{
		var count [26]int
		for _, r:= range str{
			count[r-'a']++
		}
		hashTable[count] = append(hashTable[count], str)
	}

	realRes := [][]string{}
	for _, res := range hashTable{
		realRes = append(realRes, res)
	}
	return realRes
}

