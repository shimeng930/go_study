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
