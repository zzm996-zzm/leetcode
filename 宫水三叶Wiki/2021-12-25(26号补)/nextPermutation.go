package main

//31. 下一个排列 https://leetcode-cn.com/problems/next-permutation/submissions/
//难度：中等
func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}

	n := len(nums)
	i, j, k := n-2, n-1, n-1

	//find
	for i >= 0 && nums[i] >= nums[j] {
		i--
		j--
	}

	if i >= 0 { // 不是最后一个排列
		// find: A[i]<A[k]
		for nums[i] >= nums[k] {
			k--
		}
		// swap A[i], A[k]
		nums[i], nums[k] = nums[k], nums[i]
	}

	for i, j := j, n-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}

	return
}
