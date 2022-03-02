package slice

type CompareFn[T any] func(left T, right T) int64

func OrderBy[T any](items []T, fn CompareFn[T]) []T {
	return quickSort(items, 0, len(items)-1, fn)
}
func quickSort[T any](arr []T, left, right int, fn CompareFn[T]) []T {
	if left < right {
		PartitionIndex := partition(arr, left, right, fn)
		quickSort(arr, left, PartitionIndex-1, fn)
		quickSort(arr, PartitionIndex+1, right, fn)
	}
	return arr
}
func partition[T any](arr []T, left, right int, fn CompareFn[T]) int {
	pivot := left
	index := pivot + 1
	for i := index; i <= right; i++ {
		if fn(arr[i], arr[pivot]) < 0 {
			swap(arr, i, index)
			index += 1
		}
	}
	swap(arr, pivot, index-1)
	return index - 1
}

func swap[T any](arr []T, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
