package slice_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wanglihui/go-func/slice"
)

func TestFind(t *testing.T) {
	t.Run("Find should be ok", func(t *testing.T) {
		arr := []string{"1", "2", "3"}
		it, ok := slice.Find(arr, func(t string, i int) bool {
			x, _ := strconv.Atoi(t)
			return x == 1
		})
		assert.Equal(t, ok, true)
		assert.Equal(t, *it, "1")
	})
}
