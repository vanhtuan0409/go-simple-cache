package cache

import (
	"container/list"
)

type evictFn func(c *cache) bool

func lruEvictFn(c *cache) bool {
	item := c.queue.Remove(c.queue.Back()).(*cacheItem)
	delete(c.items, item.key)
	return true
}

func lfuEvictFn(c *cache) bool {
	it := c.queue.Front()
	var evictedElement *list.Element
	var evictedItem *cacheItem
	for it != nil {
		item := it.Value.(*cacheItem)
		if evictedItem == nil || item.useCount <= evictedItem.useCount {
			evictedElement = it
			evictedItem = item
		}
		it = it.Next()
	}
	c.queue.Remove(evictedElement)
	delete(c.items, evictedItem.key)
	return true
}
