package main

//997. 找到小镇的法官  https://leetcode-cn.com/problems/find-the-town-judge/
//难度：简单
func findJudge(n int, trust [][]int) int {
	inDegrees := make([]int, n+1)
	outDegrees := make([]int, n+1)
	for _, t := range trust {
		inDegrees[t[1]]++
		outDegrees[t[0]]++
	}
	for i := 1; i <= n; i++ {
		if inDegrees[i] == n-1 && outDegrees[i] == 0 {
			return i
		}
	}
	return -1
}
