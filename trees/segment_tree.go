package trees

type SegmentTree[T any] interface {
	Query(lo, hi int) T
	Update(pos int, val T)
	SetInitQueryValue(initQueryValue T)
}

type segmentTreeWithArray[T any] struct {
	data             []T
	size             int
	initQueryValue   T
	combineRangeFunc func(T, T) T
}

func NewSegmentTreeWithArray[T any](
	size int, combineRangeFunc func(T, T) T,
) SegmentTree[T] {
	return &segmentTreeWithArray[T]{
		data:             make([]T, size*2),
		size:             size,
		combineRangeFunc: combineRangeFunc,
	}
}

func (st *segmentTreeWithArray[T]) SetInitQueryValue(initQueryValue T) {
	st.initQueryValue = initQueryValue
}

func (st *segmentTreeWithArray[T]) Query(lo, hi int) T {
	lo += st.size
	hi += st.size

	res := st.initQueryValue
	for lo < hi {
		if lo&1 != 0 {
			res = st.combineRangeFunc(res, st.data[lo])
			lo++
		}
		if hi&1 != 0 {
			hi--
			res = st.combineRangeFunc(res, st.data[hi])
		}
		lo >>= 1
		hi >>= 1
	}
	return res
}

func (st *segmentTreeWithArray[T]) Update(pos int, val T) {
	pos += st.size
	st.data[pos] = val

	for pos > 1 {
		pos >>= 1
		st.data[pos] = st.combineRangeFunc(st.data[2*pos], st.data[2*pos+1])
	}
}

type segmentTreeWithMap[T any] struct {
	data             map[int]T
	size             int
	initQueryValue   T
	combineRangeFunc func(T, T) T
}

func NewSegmentTreeWithMap[T any](
	size int, combineRangeFunc func(T, T) T,
) SegmentTree[T] {
	return &segmentTreeWithMap[T]{
		data:             make(map[int]T),
		size:             size,
		combineRangeFunc: combineRangeFunc,
	}
}

func (st *segmentTreeWithMap[T]) SetInitQueryValue(initQueryValue T) {
	st.initQueryValue = initQueryValue
}

func (st *segmentTreeWithMap[T]) Query(lo, hi int) T {
	lo += st.size
	hi += st.size

	res := st.initQueryValue
	for lo < hi {
		if lo&1 != 0 {
			res = st.combineRangeFunc(res, st.data[lo])
			lo++
		}
		if hi&1 != 0 {
			hi--
			res = st.combineRangeFunc(res, st.data[hi])
		}
		lo >>= 1
		hi >>= 1
	}
	return res
}

func (st *segmentTreeWithMap[T]) Update(pos int, val T) {
	pos += st.size
	st.data[pos] = val

	for pos > 1 {
		pos >>= 1
		st.data[pos] = st.combineRangeFunc(st.data[2*pos], st.data[2*pos+1])
	}
}
