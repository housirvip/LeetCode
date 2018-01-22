package leetcode

func longestPalindrome(str string) string {
	start, end, len1, len2, lenPalindrome := 0, 0, 0, 0, 0
	for i := 0; i < len(str); i++ {
		len1 = palindromeLen(str, i, i)
		len2 = palindromeLen(str, i, i+1)
		lenPalindrome = maxInt(len1, len2)
		if lenPalindrome > end-start {
			start, end = i-(lenPalindrome-1)/2, i+lenPalindrome/2
		}
	}
	return str[start:end+1]
}

func maxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func palindromeLen(str string, left, right int) int {
	same := 0
	for i := 0; left-i >= 0 && right+i < len(str); i++ {
		if str[left-i] == str[right+i] {
			same++
		} else {
			break
		}
	}
	if left == right {
		return 2*same - 1
	} else {
		return 2 * same
	}
}

// 翻转字符串，消耗的时间只有n/2
func reverseString(source string) string {
	runes := []rune(source)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}
