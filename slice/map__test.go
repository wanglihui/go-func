package slice_test

import (
	"errors"
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wanglihui/go-func/slice"
)

func TestMap(t *testing.T) {
	t.Run("Map type int should be ok", func(t *testing.T) {
		arr := []int{2, 3, 4, 5}
		retArr := slice.Map(arr, func(item int, idx int) int64 {
			return int64(item)
		})
		assert.Equal(t, len(arr), len(retArr))

		retArr2 := slice.Map(arr, func(item int, idx int) string {
			s := strconv.Itoa(item)
			s += "-"
			return s
		})
		assert.Equal(t, len(retArr2), len(arr))
		fmt.Println(retArr2)
	})
	t.Run("map type string should be ok", func(t *testing.T) {
		arr := []string{"1", "2", "3"}
		retArr := slice.Map(arr, func(item string, idx int) int64 {
			a, err := strconv.Atoi(item)
			if err != nil {
				t.Error(err)
				t.FailNow()
			}
			return int64(a)
		})
		assert.Equal(t, len(arr), len(retArr))
	})

	t.Run("mapasync should be ok", func(t *testing.T) {
		arr := []string{"1", "2", "3"}
		retArr, err := slice.MapAsync(arr, func(t string, i int) (int64, error) {
			if a, err := strconv.Atoi(t); err != nil {
				return 0, err
			} else {
				return int64(a), nil
			}
		})
		assert.Equal(t, len(arr), len(retArr))
		assert.NoError(t, err)
	})
	t.Run("mapasync should throw error", func(t *testing.T) {
		arr := []string{"1", "2", "3"}
		_, err := slice.MapAsync(arr, func(t string, i int) (int64, error) {
			return 0, errors.New("test error")
		})
		assert.Error(t, err)
	})
}
