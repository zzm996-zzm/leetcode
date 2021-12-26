package easycode

import "testing"

func TestCheckList(t *testing.T) {
	for _, one := range []struct {
		List []int
		Ok   bool
		Next int
	}{
		{[]int{3, 5, 7, 9, 11}, true, 13},
		{[]int{2, 4, 8, 16, 32}, true, 64},
		{[]int{1, 2, 4, 8, 16}, true, 32},
		{[]int{2, 15, 41, 80}, true, 132},
		{[]int{1, 2, 6, 15, 31}, true, 56},
		{[]int{1, 1, 3, 15, 105, 945}, true, 10395},
		{[]int{1, 1, 3, 0, 105, 945}, false, 0},
		{[]int{2, 14, 64, 202, 502, 1062, 2004}, true, 3474},
		{[]int{2, 6, 12, 20, 30}, true, 42},
		{[]int{2, 5, 10, 17}, true, 26},
		{[]int{11, 9, 7, 5, 3}, true, 1},
		{[]int{32, 16, 8, 4, 2}, false, 0},
		{[]int{2, -4, 8, -16, 32}, true, -64},
		{[]int{-3, 5, -1, 9, 7, 25}, true, 39},
		{[]int{0}, false, 0},
		{[]int{0, 1, 4}, true, 9},
		{[]int{0, 0, 0, 0}, true, 0},
		{[]int{5, 5, 5, 5}, true, 5},
		{[]int{1, 2, 2, 4}, true, 4},
		{[]int{1, 2, 2, 2, 3, 3, 3, 4, 5, 0}, true, 16},
		{[]int{1, 2, 2, 3, 7}, true, 16},
		{[]int{1, 1, 2, 2, 3, 4, 5, 6}, false, 0},
		{[]int{1, 0, 1, 0, 1, 0, 1, 0}, true, 1},
		{[]int{1, 0, 2, 3, 1, 2, 2, 1000, 2000, 5, 1, 0, 1, 0, 1, 0, 2, 3, 1, 2, 2, 1000, 2000, 5, 1, 0, 1, 0, 1, 0, 2, 3, 1, 2, 2, 1000, 2000, 5, 1, 0, 1, 0, 1, 0, 2, 3, 1, 2, 2, 1000, 2000, 5, 1, 0, 1, 0, 1, 0, 2, 3, 1, 2, 2, 1000, 2000, 5, 1, 0, 1, 0}, false, 0},
	} {
		next, ok := checkList(one.List)
		if ok == one.Ok && next == one.Next {
			t.Log("correct", "input", one.List, "get", next, ok, "want", one.Next, one.Ok)
		} else {
			t.Error("wrong", "input", one.List, "get", next, ok, "want", one.Next, one.Ok)
		}
	}
}
