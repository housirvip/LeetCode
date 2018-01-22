package leetcode

// 复杂度o（n3）
func lengthOfLongestSubstring(source string) int {
	len := len(source)
	res := 0
	if len > 0 {
		res = 1
	}
	for i := 0; i < len; i++ {
		for j := i + 1; j < len; j++ {
			if lastRepeat(source, i, j) {
				break
			} else {
				temp := j - i + 1
				if temp > res {
					res = temp
				}
			}
		}
	}
	return res
}

func lastRepeat(source string, start, end int) bool {
	for i := start; i < end; i++ {
		if source[i] == source[end] {
			return true
		}
	}
	return false
}
