package slice

import (
	"golang.org/x/sync/errgroup"
)

func Map[T any, V any](arr []T, f func(T, int) V) []V {
	result := make([]V, len(arr))
	for i, elem := range arr {
		result[i] = f(elem, i)
	}
	return result
}

func MapAsync[T any, V any](arr []T, f func(T, int) (V, error)) ([]V, error) {
	rets := make([]V, len(arr))
	group := new(errgroup.Group)
	for i, t := range arr {
		i, t := i, t
		group.Go(func() error {
			if v, err := f(t, i); err != nil {
				return err
			} else {
				rets[i] = v
				return nil
			}
		})
	}
	if err := group.Wait(); err != nil {
		return nil, err
	}
	return rets, nil
}
