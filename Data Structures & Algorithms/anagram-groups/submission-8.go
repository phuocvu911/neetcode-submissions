
func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0{
		return nil
	}

	resTable := make(map[string][]string)

	for _, str := range strs{
		sorted := sortString(str)
		resTable[sorted] = append(resTable[sorted], str)
	}

	realRes := [][]string{}
	for _, res := range resTable{
		realRes = append(realRes, res)
	}
	return realRes

}

func sortString(s string) string{
	chars := []rune(s)
	sort.Slice(chars, func(i,j int) bool{
		return chars[i] < chars[j]
	})
	return string(chars)
}
