package main

// 66. 加一 https://leetcode-cn.com/problems/plus-one/
// 难度：简单
// 依稀记得这道题两个月前写 写了50多行 真的难看
func plusOne(digits []int) []int {
	inc := 1
	for i := len(digits) - 1; i >= 0 && inc > 0; i-- {
		ans := digits[i] + inc
		inc = ans / 10
		digits[i] = ans % 10
	}
	if inc > 0 {
		digits = append(digits, 1)
		digits[0], digits[len(digits)-1] = digits[len(digits)-1], digits[0]
	}
	return digits
}
