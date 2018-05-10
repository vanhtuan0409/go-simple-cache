package cache

func Example() {
	lruCache := NewLRUCache(2)
	lruCache.Set("x", 1)
	lruCache.Get("x")
}
