package leetcode

import "testing"

func Test_longestPalindrome(t *testing.T) {
	//t.Log(reverseString("abcdasdfghjkldcba"))
	t.Log(longestPalindrome("babad"))
	t.Log(longestPalindrome("abcacbadf"))
	t.Log(longestPalindrome("cbbca"))
	t.Log(longestPalindrome("cbbd"))
}

func Test_Manacher(t *testing.T) {
	str := "abcacbadf"
	t.Log(changeStr(str))
	t.Log(manacher(str))
	t.Log(manacher("babad"))
	t.Log(manacher("cbbca"))
}
