package slice

func UniqBy[T any, V comparable](arr []T, identity Identity[T, V]) []T {
	retArr := make([]T, 0, len(arr))
	for _, item := range arr {
		hasContain := false
		for _, item2 := range retArr {
			if identity(item) == identity(item2) {
				hasContain = true
				break
			}
		}
		if !hasContain {
			retArr = append(retArr, item)
		}
	}
	return retArr
}

func Uniq[T comparable](arr []T) []T {
	return UniqBy(arr, func(t T) T {
		return t
	})
}
