func isValid(s string) bool {
	hash := map[rune]rune{
		'[': ']',
		'{': '}',
		'(': ')',
	}
	stack := []rune{}
	for _, char := range s {
		if val, ok := hash[char]; ok {
			stack = append(stack, val)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != char {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
