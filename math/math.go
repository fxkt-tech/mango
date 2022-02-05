package math

import "constraints"

// 整型绝对值
func Abs[T constraints.Integer | constraints.Float](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func MaxMin[T constraints.Integer | constraints.Float](r, g, b T) (T, T) {
	var max, min T
	if r > g {
		max, min = r, g
	} else {
		max, min = g, r
	}
	if b > max {
		max = b
	} else if b < min {
		min = b
	}
	return max, min
}

func Max[T constraints.Integer | constraints.Float](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Min[T constraints.Integer | constraints.Float](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func MaxInSlice[T constraints.Integer | constraints.Float](s ...T) T {
	slen := len(s)
	switch slen {
	case 0:
		return 0
	case 1:
		return s[0]
	default:
		max := s[0]
		for i := 1; i < len(s); i++ {
			if s[i] > max {
				max = s[i]
			}
		}
		return max
	}
}

func Clip[T constraints.Integer | constraints.Float](x, a, b T) T {
	return Min(Max(x, a), b)
}
