package main

// https://leetcode-cn.com/problems/zigzag-conversion/submissions/

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	period := 2*numRows - 2

	strs := make([][]byte, numRows)
	//初始化 strs

	for i := 0; i < numRows; i++ {
		strs[i] = make([]byte, 0)
	}

	for i := 0; i < len(s); i++ {
		mod := i % period
		if mod < numRows {
			strs[mod] = append(strs[mod], s[i])
		} else {
			strs[period-mod] = append(strs[period-mod], s[i])
		}
	}

	ans := ""

	for _, v := range strs {
		ans += string(v)
	}

	return ans
}
