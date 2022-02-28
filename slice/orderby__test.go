package slice_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wanglihui/go-funk/slice"
)

func TestOrderBy(t *testing.T) {
	t.Run("OrderBy should be ok", func(t *testing.T) {
		arr := []string{"3", "1", "2"}
		newArr := slice.OrderBy(arr, func(left string, right string) int64 {
			lefti, _ := strconv.Atoi(left)
			righti, _ := strconv.Atoi(right)
			return int64(lefti) - int64(righti)
		})
		assert.Equal(t, len(arr), len(newArr))
		assert.Equal(t, "2", newArr[1])
	})
}
