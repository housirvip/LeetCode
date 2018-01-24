package leetcode

func twoSum(nums []int, target int) []int {
	index :=[]int{0,0}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i]+nums[j] == target {
				index[0] = j
				index[1] = i
				return index
			}
		}
	}
	return index
}
