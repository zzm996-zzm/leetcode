package main

// 1518. 换酒问题 https://leetcode-cn.com/problems/water-bottles/
// 难度：简单
func numWaterBottles(numBottles int, numExchange int) int {
	ans := numBottles
	for numBottles >= numExchange {
		a := numBottles / numExchange
		b := numBottles % numExchange
		ans += a
		numBottles = a + b
	}

	return ans
}
