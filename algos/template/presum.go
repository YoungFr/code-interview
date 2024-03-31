package template

// 二维前缀和模板
type NumMatrix struct {
	ps [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	var (
		m = len(matrix)
		n = len(matrix[0])
	)
	ps := make([][]int, m+1)
	for i := 0; i < len(ps); i++ {
		ps[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ps[i+1][j+1] = ps[i+1][j] + ps[i][j+1] - ps[i][j] + matrix[i][j]
		}
	}
	return NumMatrix{ps}
}

func (nm *NumMatrix) SumRegion(r1 int, c1 int, r2 int, c2 int) int {
	return nm.ps[r2+1][c2+1] - nm.ps[r2+1][c1] - nm.ps[r1][c2+1] + nm.ps[r1][c1]
}
