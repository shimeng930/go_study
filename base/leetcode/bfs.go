package leetcode

import "fmt"

// 将一层的数据放到队列里
// 取当前队列长度为进行一次遍历，这一次的出队的数据都在同一层
// 按照先进先出的顺序，出一个节点数据，然后将这个节点附近关联的下一层都加入队列
// 继续出下一个节点，添加关联数据
func bfsTree(root *TreeNode) {
	var arr = []*TreeNode{root}
	var level = 1
	for len(arr) > 0 {
		// 当前层的数量
		size := len(arr)
		fmt.Print("level", level, "--->data:")
		for i := 0; i < size; i++ {
			top := arr[i]
			if top.Left != nil {
				arr = append(arr, top.Left)
			}
			if top.Right != nil {
				arr = append(arr, top.Right)
			}
			fmt.Print(top.Val, ",")
		}
		arr = arr[size:]
		level++
		fmt.Println()
	}
}

func bfsTree2(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	var arr = []*TreeNode{root}
	var level = 0

	for len(arr) > 0 {
		//res = append(res, []int{})
		size := len(arr)
		var levelRes []int
		for i := 0; i < size; i++ {
			// res[level] = append(res[level], arr[i].Val)
			levelRes = append(levelRes, arr[i].Val)
			top := arr[i]
			if top.Left != nil {
				arr = append(arr, top.Left)
			}
			if top.Right != nil {
				arr = append(arr, top.Right)
			}
		}
		res = append(res, levelRes)
		level++
		arr = arr[size:]
	}

	return res
}

// lc1905腐烂的橘子
func orangesRotting(grid [][]int) int {
	var refresh, res int
	var rots [][2]int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 2 {
				rots = append(rots, [2]int{i, j})
			} else if grid[i][j] == 1 {
				refresh++
			}
		}
	}

	for len(rots) > 0 {
		if refresh == 0 {
			return res
		}
		res++

		size := len(rots)
		for i := 0; i < size; i++ {
			x, y := rots[i][0], rots[i][1]
			if rot(x+1, y, grid) {
				refresh--
				rots = append(rots, [2]int{x + 1, y})
			}
			if rot(x-1, y, grid) {
				refresh--
				rots = append(rots, [2]int{x - 1, y})
			}
			if rot(x, y+1, grid) {
				refresh--
				rots = append(rots, [2]int{x, y + 1})
			}
			if rot(x, y-1, grid) {
				refresh--
				rots = append(rots, [2]int{x, y - 1})
			}
		}
		rots = rots[size:]
	}
	if refresh == 0 {
		return res
	} else {
		return -1
	}
}

func rot(i, j int, grid [][]int) bool {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[i]) {
		return false
	}
	if grid[i][j] == 1 {
		grid[i][j] = 2
		return true
	}
	return false
}
