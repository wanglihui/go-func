package slice

// Chunk
//
// As a example:
//     arr := []int{1,2,3}
//     newArr := Chunk(arr, 2)
func Chunk[T any](arr []T, size int) [][]T {
	retArr := make([][]T, 0, len(arr)/size+1)
	for idx, item := range arr {
		arrIndex := idx / size
		if idx%size == 0 {
			newArr := make([]T, 0, size)
			retArr = append(retArr, newArr)
		}
		retArr[arrIndex] = append(retArr[arrIndex], item)
	}
	return retArr
}
