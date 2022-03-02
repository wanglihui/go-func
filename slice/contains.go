package slice

// Contains
//
//    arr := []int{1,2,3}
//    b := Contains(arr, 1)
//    Output: true
func Contains[T any](arr []T, f func(t T, i int) bool) bool {
	_, ok := Find(arr, f)
	return ok
}
