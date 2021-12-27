package main

import "strconv"

// 43. 字符串相乘 https://leetcode-cn.com/problems/multiply-strings/
// 难度：中等
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	ret := ""
	for i := len(num2) - 1; i >= 0; i-- {
		ans := ""
		//拼接0
		for j := len(num2) - 1; j > i; j-- {
			ans += "0"
		}

		n2 := int(num2[i] - '0')
		inc := 0
		for k := len(num1) - 1; k >= 0 || inc > 0; k-- {
			n1 := 0
			if k >= 0 {
				n1 = int(num1[k] - '0')
			}
			temp := n1*n2 + inc
			ans = strconv.Itoa(temp%10) + ans
			inc = temp / 10
		}

		ret = addStrings(ret, ans)
	}

	return ret
}

func addStrings(num1 string, num2 string) string {
	add := 0
	ans := ""
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0 || add != 0; i, j = i-1, j-1 {
		var x, y int
		if i >= 0 {
			x = int(num1[i] - '0')
		}
		if j >= 0 {
			y = int(num2[j] - '0')
		}
		result := x + y + add
		ans = strconv.Itoa(result%10) + ans
		add = result / 10
	}
	return ans
}
