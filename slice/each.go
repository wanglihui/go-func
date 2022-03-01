package slice

import "golang.org/x/sync/errgroup"

func Each[T any](arr []T, f func(it T, idx int)) {
	for idx, it := range arr {
		f(it, idx)
	}
}

func EachAsync[T any](arr []T, f func(it T, idx int) error) error {
	group := new(errgroup.Group)
	for idx, it := range arr {
		idx, it := idx, it
		group.Go(func() error {
			return f(it, idx)
		})
	}
	return group.Wait()
}
