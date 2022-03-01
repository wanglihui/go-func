package slice_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wanglihui/go-func/slice"
)

func TestChunk(t *testing.T) {
	t.Run("chunk should be ok", func(t *testing.T) {
		arr := []int{1, 2, 3}
		retArr := slice.Chunk(arr, 2)
		assert.Equal(t, len(retArr), 2)
		assert.Equal(t, len(retArr[1]), 1)
	})

	t.Run("chunk should be ok", func(t *testing.T) {
		arr := []int{1, 2, 3, 4}
		retArr := slice.Chunk(arr, 2)
		assert.Equal(t, len(retArr), 2)
		assert.Equal(t, len(retArr[1]), 2)
	})
}
