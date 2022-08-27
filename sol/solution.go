package sol

func isValid(s string) bool {
	if len(s) <= 1 {
		return false
	}
	stack := []rune{}
	popMap := map[rune]rune{')': '(', ']': '[', '}': '{'}
	for _, val := range s {
		if popVal, exist := popMap[val]; !exist {
			stack = append(stack, val)
		} else {
			if len(stack) == 0 {
				return false
			}
			topVal := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if topVal != popVal {
				return false
			}
		}
	}
	return len(stack) == 0
}
