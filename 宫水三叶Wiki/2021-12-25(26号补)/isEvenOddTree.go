package main

// 1609. 奇偶树  https://leetcode-cn.com/problems/even-odd-tree/
// 难度：中等

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//DFS
func isEvenOddTree2(root *TreeNode) bool {
	m := make(map[int]int)
	return dfs(root, 0, m)
}

func dfs(node *TreeNode, level int, m map[int]int) bool {
	flag := level%2 == 0
	cur := node.Val
	prev := 0
	if v, has := m[level]; has {
		prev = v
	} else if flag {
		prev = 0
	} else {
		prev = 0xFFFFFFF
	}
	if flag && (cur%2 == 0 || cur <= prev) {
		return false
	}
	if !flag && (cur%2 != 0 || cur >= prev) {
		return false
	}

	m[level] = cur

	if node.Left != nil && !dfs(node.Left, level+1, m) {
		return false
	}
	if node.Right != nil && !dfs(node.Right, level+1, m) {
		return false
	}

	return true
}

//BFS
func isEvenOddTree(root *TreeNode) bool {
	if root == nil {
		return false
	}
	queue := New()
	flag := true
	prev := 0
	queue.Enqueue(root)
	for !queue.IsEmpty() {
		Len := queue.Size()
		if flag {
			prev = 0
		} else {
			prev = 0xFFFFFFF
		}
		for i := 0; i < Len; i++ {
			node := queue.Dequeue().(*TreeNode)
			cur := node.Val
			if flag && (cur%2 == 0 || cur <= prev) {
				return false
			}
			if !flag && (cur%2 != 0 || cur >= prev) {
				return false
			}
			prev = cur

			if node.Left != nil {
				queue.Enqueue(node.Left)
			}
			if node.Right != nil {
				queue.Enqueue(node.Right)
			}
		}
		flag = !flag
	}

	return true
}

//首先实现一个队列....
type Queue struct {
	arr []interface{}
}

func New() *Queue {
	return &Queue{arr: make([]interface{}, 0, 10)}
}

func (q *Queue) Enqueue(x interface{}) {
	q.arr = append(q.arr, x)
}
func (q *Queue) Dequeue() interface{} {
	if !q.IsEmpty() {
		res := q.arr[0]
		q.arr = q.arr[1:]
		return res
	}

	return -1
}

func (q *Queue) Front() interface{} {
	return q.arr[0]
}
func (q *Queue) IsEmpty() bool {
	return len(q.arr) == 0
}
func (q *Queue) Size() int {
	return len(q.arr)
}
