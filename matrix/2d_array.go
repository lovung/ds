package matrix

func New2D[T any](m, n int, defaultVal ...T) [][]T {
	data := make([][]T, m)
	for i := range data {
		data[i] = make([]T, n)
		if len(defaultVal) > 0 {
			for j := range data[i] {
				data[i][j] = defaultVal[0]
			}
		}
	}
	return data
}

func At[T any](mat [][]T, i, j int) T {
	var zero T
	if !InBound(mat, i, j) {
		return zero
	}
	return mat[i][j]
}

func InBound[T any](mat [][]T, i, j int) bool {
	if i < 0 || i >= len(mat) {
		return false
	}
	if len(mat) > 0 && (j < 0 || j >= len(mat[0])) {
		return false
	}
	return true
}
func UpdateAt[T any](mat [][]T, i, j int, val T) {
	if !InBound(mat, i, j) {
		return
	}
	mat[i][j] = val
}

func ForEach[T any](mat [][]T, f func(i, j int)) {
	for i := range mat {
		for j := range mat[i] {
			f(i, j)
		}
	}
}
