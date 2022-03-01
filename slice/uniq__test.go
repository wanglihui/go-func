package slice_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wanglihui/go-func/slice"
)

func TestUniqBy(t *testing.T) {
	t.Run("uniqby should be ok", func(t *testing.T) {
		arr := []int{1, 2, 1, 3, 1, 2}
		retArr := slice.UniqBy(arr, func(t int) int {
			return t
		})
		assert.Equal(t, len(retArr), 3)
	})
	t.Run("uniqby should be ok", func(t *testing.T) {
		type user struct {
			ID       int
			Username string
		}
		arr := []user{
			{
				ID:       1,
				Username: "a1",
			},
			{
				ID:       2,
				Username: "a2",
			},
			{
				ID:       1,
				Username: "a2",
			},
		}
		retArr := slice.UniqBy(arr, func(t user) int {
			return t.ID
		})
		assert.Equal(t, len(retArr), 2)
	})
}

func TestUiqu(t *testing.T) {
	t.Run("uniq should be ok", func(t *testing.T) {
		arr := []int{1, 2, 3, 2, 3, 4}
		retArr := slice.Uniq(arr)
		assert.Equal(t, len(retArr), 4)
	})
}
