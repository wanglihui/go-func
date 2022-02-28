package slice_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wanglihui/go-funk/slice"
)

func TestDiff(t *testing.T) {
	t.Run("diff should be ok", func(t *testing.T) {
		a1 := []string{"1", "3", "5"}
		a2 := []string{"5", "7", "9"}
		diff := slice.Diff(a1, a2, func(a, b string) bool {
			return a == b
		})
		assert.Equal(t, len(diff), 4)
		assert.Equal(t, diff[0], a1[0])
	})
}
