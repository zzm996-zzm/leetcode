package main

// 58. 最后一个单词的长度 https://leetcode-cn.com/problems/length-of-last-word/
// 难度：简单

func lengthOfLastWord(s string) int {
	count := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			count++
		} else if s[i] == ' ' && count != 0 {
			break
		}
	}

	return count
}
