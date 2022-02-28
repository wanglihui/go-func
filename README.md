go-func
---

don't to create repeat wheel

### function index

- Map
- MapAsync
- Find
- OrderBy
- Filter
- Contains
-

### usage

```
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
```