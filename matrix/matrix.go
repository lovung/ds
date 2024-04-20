package matrix

type Matrix[T any] struct {
	data       [][]T
	defaultVal T
}

var Dirs = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

type CellFn[T any] func(val T, i, j int)

func NewFrom[S, D any](source [][]S, defaultVal ...D) *Matrix[D] {
	data := make([][]D, len(source))
	for i := range data {
		data[i] = make([]D, len(source[i]))
		if len(defaultVal) > 0 {
			for j := range data[i] {
				data[i][j] = defaultVal[0]
			}
		}
	}
	mat := &Matrix[D]{
		data: data,
	}
	if len(defaultVal) > 0 {
		mat.defaultVal = defaultVal[0]
	}
	return mat
}

// New m row, n col matrix
func New[T any](m, n int, defaultVal ...T) *Matrix[T] {
	data := make([][]T, m)
	for i := range data {
		data[i] = make([]T, n)
		if len(defaultVal) > 0 {
			for j := range data[i] {
				data[i][j] = defaultVal[0]
			}
		}
	}
	mat := &Matrix[T]{
		data: data,
	}

	if len(defaultVal) > 0 {
		mat.defaultVal = defaultVal[0]
	}
	return mat
}

func NewSquare[T any](n int, defaultVal ...T) *Matrix[T] {
	data := make([][]T, n)
	for i := range data {
		data[i] = make([]T, n)
		if len(defaultVal) > 0 {
			for j := range data[i] {
				data[i][j] = defaultVal[0]
			}
		}
	}
	mat := &Matrix[T]{
		data: data,
	}

	if len(defaultVal) > 0 {
		mat.defaultVal = defaultVal[0]
	}
	return mat
}

func (mat *Matrix[T]) At(i, j int) T {
	if !mat.InBound(i, j) {
		return mat.defaultVal
	}
	return mat.data[i][j]
}

func (mat *Matrix[T]) InBound(i, j int) bool {
	return In(mat.data, i, j)
}

func (mat *Matrix[T]) UpdateAt(i, j int, val T) {
	if mat.InBound(i, j) {
		mat.data[i][j] = val
	}
}

func (mat *Matrix[T]) ForEach(f CellFn[T]) {
	for i := range mat.data {
		for j := range mat.data[i] {
			f(mat.At(i, j), i, j)
		}
	}
}

func (mat *Matrix[T]) ForEachNearBy(i, j int, f CellFn[T]) {
	ForEachNearBy(mat.data, i, j, f)
}

func (mat *Matrix[T]) Data() [][]T {
	return mat.data
}

func In[T any](mat [][]T, i, j int) bool {
	if i < 0 || j < 0 || i >= len(mat) || j >= len(mat[0]) {
		return false
	}
	return true
}

func ForEachNearBy[T any](mat [][]T, i, j int, f CellFn[T]) {
	for _, d := range Dirs {
		x, y := i+d[0], j+d[1]
		if In(mat, x, y) {
			f(mat[x][y], x, y)
		}
	}
}
