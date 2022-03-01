package slice_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wanglihui/go-func/slice"
)

func TestEach(t *testing.T) {
	t.Run("each should be ok", func(t *testing.T) {
		arr := []int{1, 2, 3}
		sum := 0
		slice.Each(arr, func(it int, _ int) {
			sum += it
		})
		assert.Equal(t, sum, 6)
	})

}

func TestEachAsync(t *testing.T) {
	t.Run("eachasync should be ok", func(t *testing.T) {
		arr := []int{1, 2, 3}
		sum := 0
		err := slice.EachAsync(arr, func(t int, i int) error {
			sum += t
			return nil
		})
		assert.Equal(t, sum, 6)
		assert.NoError(t, err)
	})

	t.Run("eachasync should return error", func(t *testing.T) {
		arr := []int{1, 2, 3}
		sum := 0
		err := slice.EachAsync(arr, func(t int, i int) error {
			return errors.New("test error")
		})
		assert.Error(t, err)
		assert.Equal(t, sum, 0)
	})
}
