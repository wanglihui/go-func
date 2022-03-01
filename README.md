go-func
---

[![Go](https://github.com/wanglihui/go-func/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/wanglihui/go-func/actions/workflows/go.yml)

don't to create repeat wheel

### function index

- Map
- MapAsync
- Find
- OrderBy
- Filter
- Contains
- Uniq
- UniqBy
- Diff
- DiffBy
- Chunk
- Each
- EachAsync

### usage

```golang
		//Map
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

		// MapAsync
		arr := []string{"1", "2", "3"}
		_, err := slice.MapAsync(arr, func(t string, i int) (int64, error) {
			return 0, errors.New("test error")
		})
		assert.Error(t, err)
```

### test

go test ./... -cover

```
ok  	github.com/wanglihui/go-func/slice	(cached)	coverage: 98.4% of statements
```