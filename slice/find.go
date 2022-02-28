package slice

func Find[T any](arr []T, f func(t T, i int) bool) (it *T, ok bool) {
	ok = false
	for i, t := range arr {
		if b := f(t, i); b {
			it = &t
			ok = true
			return
		}
	}
	return nil, false
}
