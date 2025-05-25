package leetcode

// lc 54 螺旋矩阵
func spiralOrder(matrix [][]int) []int {
	var res []int

	var l, r, top, bottom = 0, len(matrix[0]) - 1, 0, len(matrix) - 1

	for l <= r && top <= bottom {
		// 左到右
		var i = top
		var j = l
		for ; j <= r; j++ {
			res = append(res, matrix[i][j])
		}

		// 上到下
		j = r
		i++
		for ; i <= bottom; i++ {
			res = append(res, matrix[i][j])
		}

		// 右到左
		j--
		if top != bottom {
			i = bottom
			for ; j >= l; j-- {
				res = append(res, matrix[i][j])
			}
		}

		// 下到上
		i--
		if l != r {
			j = l
			for ; i > top; i-- {
				res = append(res, matrix[i][j])
			}
		}
		l++
		r--
		top++
		bottom--
	}
	return res
}

func setZeroes(matrix [][]int) {
	var m, n = len(matrix), len(matrix[0])
	var row0 bool
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				if i == 0 {
					row0 = true
				} else {
					matrix[i][0] = 0
				}
				matrix[0][j] = 0
			}
		}
	}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == 0 {
				if row0 {
					matrix[i][j] = 0
				}
			} else if matrix[i][0] == 0 {
				matrix[i][j] = 0
			}

			if matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
}

func rotateMatrix(matrix [][]int) {
	n := len(matrix)
	// 置换 [i, j] -> [n-j, n-i]
	for i := 0; i < n; i++ {
		for j := 0; j < n-i; j++ {
			matrix[i][j], matrix[n-1-j][n-1-i] = matrix[n-1-j][n-1-i], matrix[i][j]
		}
	}

	// 翻转
	for i := 0; i < n/2; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j], matrix[n-1-i][j] = matrix[n-1-i][j], matrix[i][j]
		}
	}
}
