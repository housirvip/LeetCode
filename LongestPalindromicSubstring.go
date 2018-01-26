package leetcode

import (
	"bytes"
)

func longestPalindrome(str string) string {
	start, end, len1, len2, lenPalindrome := 0, 0, 0, 0, 0
	for i := 0; i < len(str); i++ {
		// 奇偶要分开讨论
		len1 = palindromeLen(str, i, i)
		len2 = palindromeLen(str, i, i+1)
		lenPalindrome = maxInt(len1, len2)
		if lenPalindrome > end-start {
			start, end = i-(lenPalindrome-1)/2, i+lenPalindrome/2
		}
	}
	return str[start:end+1]
}

// golang中没有max，min这类函数，自己写一个
func maxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func palindromeLen(str string, left, right int) int {
	i := 0
	for i = 0; left-i >= 0 && right+i < len(str); i++ {
		if str[left-i] != str[right+i] {
			break
		}
	}
	return 2*i - 1 + right - left
}

// 翻转字符串，消耗的时间只有n/2，这个方法并没用到
func reverseString(source string) string {
	runes := []rune(source)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}

// 著名的马拉车算法，将复杂度回归线性O(n)
func manacher(source string) string {
	// 先对字符串增加#字符
	str := changeStr(source)
	lenStr := len(str)
	prGroup := make([]int, lenStr)
	tempCentre, tempRight, prMax, prMaxIndex := 0, 0, 0, 1
	for i := 1; i < len(str)-1; i++ {
		// golang没有三元运算符
		if i <= tempRight {
			prGroup[i] = minInt(tempRight-i, prGroup[2*tempCentre-i])
		} else {
			prGroup[i] = 0
		}
		for str[i-prGroup[i]-1] == str[i+prGroup[i]+1] {
			prGroup[i]++
		}
		if tempRight < prGroup[i]+i {
			tempCentre = i
			tempRight = i + prGroup[i]
		}
		if prMax < prGroup[i] {
			prMax = prGroup[i]
			prMaxIndex = i
		}
	}
	return source[(prMaxIndex-1-prMax)/2:(prMaxIndex-1+prMax)/2]
}

func minInt(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func changeStr(source string) string {
	b := bytes.Buffer{}
	b.WriteString("$#")
	for i := 0; i < len(source); i++ {
		b.WriteString(string(source[i]))
		b.WriteString("#")
	}
	b.WriteString("&")
	return b.String()
}
