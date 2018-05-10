package cache

type cacheItem struct {
	key      string
	value    interface{}
	useCount int
}

func (i *cacheItem) increaseUseCount() {
	i.useCount++
}
