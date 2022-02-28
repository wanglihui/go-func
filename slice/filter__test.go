package slice_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wanglihui/go-funk/slice"
)

func TestFilter(t *testing.T) {
	t.Run("filter should be ok", func(t *testing.T) {
		arr := []int{1, 2, 3}
		newArr := slice.Filter(arr, func(t int, i int) bool {
			return t > arr[0]
		})
		assert.Equal(t, len(newArr), 2)
		assert.Equal(t, newArr[0], 2)
	})
}
