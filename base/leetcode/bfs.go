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

func canFinish(numCourses int, prerequisites [][]int) bool {
	//var prem, next, courses = make(map[int]int), make(map[int][]int), make(map[int]bool)
	//for _, n := range prerequisites {
	//	prem[n[0]]++
	//	next[n[1]] = append(next[n[1]], n[0])
	//	courses[n[0]] = true
	//	courses[n[1]] = true
	//}
	//var q []int
	//for k, _ := range courses {
	//	if prem[k] == 0 {
	//		q = append(q, k)
	//	}
	//}
	//var size, learn = len(q), 0
	//for size > 0 {
	//	learn += size
	//	for i := 0; i < size; i++ {
	//		for _, n := range next[q[i]] {
	//			prem[n]--
	//			if prem[n] == 0 {
	//				q = append(q, n)
	//			}
	//		}
	//	}
	//	q = q[size:]
	//	size = len(q)
	//}
	//return learn == len(courses)

	adj := map[int][]int{}
	inDegree := make([]int, numCourses)
	for _, v := range prerequisites {
		inDegree[v[0]]++
		if _, ok := adj[v[1]]; ok {
			adj[v[1]] = append(adj[v[1]], v[0])
		} else {
			adj[v[1]] = []int{v[0]}
		}
	}
	var que []int
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			que = append(que, i)
		}
	}
	cnt := 0
	for len(que) > 0 {
		v := que[0]
		que = que[1:]
		cnt++
		for _, nei := range adj[v] {
			inDegree[nei]--
			if inDegree[nei] == 0 {
				que = append(que, nei)
			}
		}
	}

	return cnt == numCourses
}
