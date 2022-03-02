package slice_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wanglihui/go-func/slice"
)

func TestGroupBy(t *testing.T) {
	arr := []float64{1.1, 1.2, 2.1, 2.2}
	var groupbyFn = func(f float64) int64 {
		return int64(f)
	}
	ret := slice.GroupBy(arr, groupbyFn)
	fmt.Println(ret)
	// Output: 1:1.1,1.2 2:2.1,2.2
	assert.Equal(t, len(ret[1]), 2)
	assert.Equal(t, len(ret[2]), 2)
}
