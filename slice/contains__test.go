package slice_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wanglihui/go-func/slice"
)

func TestContains(t *testing.T) {
	t.Run("contains should be ok", func(t *testing.T) {
		arr := []string{"1", "2", "3"}
		b := slice.Contains(arr, func(t string, i int) bool {
			return t == "2"
		})
		assert.Equal(t, b, true)
	})
}
