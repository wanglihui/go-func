package slice

func DiffBy[T comparable](arr, arr2 []T, f func(a, b T) bool) []T {
	diffArr := make([]T, 0)
	for _, t := range arr {
		b := false
		for _, v := range arr2 {
			if b = f(t, v); b {
				break
			}
		}
		if !b {
			diffArr = append(diffArr, t)
		}
	}
	for _, v := range arr2 {
		if Contains(diffArr, func(it T, i int) bool {
			return f(it, v)
		}) {
			continue
		}
		b := false
		for _, t := range arr {
			if b = f(t, v); b {
				break
			}
		}
		if !b {
			diffArr = append(diffArr, v)
		}
	}
	return diffArr
}

func Diff[T comparable](arr1 []T, arr2 []T) []T {
	return DiffBy(arr1, arr2, func(a, b T) bool {
		return a == b
	})
}
