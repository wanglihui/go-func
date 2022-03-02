package slice

// DiffBy
//
//		a1 := []string{"1", "3", "5"}
//		a2 := []string{"5", "7", "9"}
//		diff := slice.DiffBy(a1, a2, func(a, b string) bool {
//			return a == b
//		})
//		//Output: ["1", "3", "7", "9"]
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

// Diff
//
//    arr := []string{"1", "3", "5"}
//    arr2 := []string{"2", "4", "5"}
//    diffArr := Diff(arr, arr2)
//    //Output: ["1", "3", "2", "4"]
//
func Diff[T comparable](arr1 []T, arr2 []T) []T {
	return DiffBy(arr1, arr2, func(a, b T) bool {
		return a == b
	})
}
