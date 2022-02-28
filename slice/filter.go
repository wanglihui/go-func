package slice

func Filter[T any](arr []T, f func(t T, i int) bool) []T {
	newArr := make([]T, 0)
	for i, it := range arr {
		if b := f(it, i); b {
			newArr = append(newArr, it)
		}
	}
	return newArr
}
