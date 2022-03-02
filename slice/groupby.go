package slice

// GroupBy
//
//		arr := []float64{1.1, 1.2, 2.1, 2.2}
//		var groupbyFn = func(f float64) int64 {
//			return int64(f)
//		}
//		ret := slice.GroupBy(arr, groupbyFn)
//		fmt.Println(ret)
//		// Output: 1:1.1,1.2 2:2.1,2.2
//
func GroupBy[T any, V comparable](arr []T, fn Identity[T, V]) map[V][]T {
	m := make(map[V][]T, 0)
	for _, item := range arr {
		v := fn(item)
		val := m[v]
		if len(val) == 0 {
			val = make([]T, 0)
		}
		val = append(val, item)
		m[v] = val
	}
	return m
}
