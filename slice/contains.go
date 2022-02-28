package slice

func Contains[T any](arr []T, f func(t T, i int) bool) bool {
	_, ok := Find(arr, f)
	return ok
}
