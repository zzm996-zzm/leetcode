package main

import "strings"

//1078. Bigram 分词 https://leetcode-cn.com/problems/occurrences-after-bigram/
//难度：简单
func findOcurrences(text, first, second string) (ans []string) {
	words := strings.Split(text, " ")
	for i := 2; i < len(words); i++ {
		if words[i-2] == first && words[i-1] == second {
			ans = append(ans, words[i])
		}
	}
	return
}
