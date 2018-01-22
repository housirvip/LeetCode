package leetcode

//时间复杂度为o(n)其中n为str长度
func getNext(model string) []int {
	len := len(model)
	// 声明动态长度为len的一个next数组
	next := make([]int, len)
	for i := 1; i < len; i++ {
		// 无非归零，归一，加一三种情况
		if next[i-1] > 0 && model[next[i-1]] == model[i] {
			next[i] = next[i-1] + 1
		} else if model[0] == model[i] {
			next[i] = 1
		}
		// 不修改表示next[i]=0
	}
	return next
}

//时间复杂度为o(m)其中m为str长度
func KMP(model, target string) bool {
	next := getNext(model)
	same := 0
	// 保证i始终不退回，j尽量少退回
	for i, j := 0, 0; i < len(target); i++ {
		if model[j] == target[i] {
			same++
			j++
		} else if j > 0 {
			// 前面一位的next值表示各种回退后此时匹配的字符数
			same = next[j-1]
			// j就回退到same的值，前j位都匹配完成了，只需要往后匹配
			j = same
		}
		if same == len(model) {
			// 此时刚好匹配完成，开始位置为i-length+1，结束位置为i
			return true
		}
	}
	return false
}
