package cache

import (
	"container/list"
	"errors"
	"sync"
)

var (
	KeyNotFound = errors.New("cache key not found")
)

type Cache interface {
	Set(key string, value interface{})
	Get(key string) (interface{}, error)
}

type cache struct {
	sync.Mutex
	size  int
	queue *list.List
	evict evictFn
	items map[string]*list.Element
}

func NewLRUCache(size int) Cache {
	return &cache{
		size:  size,
		queue: list.New(),
		items: map[string]*list.Element{},
		evict: lruEvictFn,
	}
}

func NewLFUCache(size int) Cache {
	return &cache{
		size:  size,
		queue: list.New(),
		items: map[string]*list.Element{},
		evict: lfuEvictFn,
	}
}

func (c *cache) Set(key string, value interface{}) {
	c.Lock()
	defer c.Unlock()

	ele, ok := c.items[key]
	if ok {
		item := ele.Value.(*cacheItem)
		item.value = value
		c.queue.MoveToFront(ele)
		return
	}

	// Handle add new item to cache
	if shouldEvictBeforeInsert(c) {
		c.evict(c)
	}
	ele = c.queue.PushFront(&cacheItem{key: key, value: value, useCount: 1})
	c.items[key] = ele
}

func (c *cache) Get(key string) (interface{}, error) {
	ele, ok := c.items[key]
	if !ok {
		return nil, KeyNotFound
	}

	c.Lock()
	defer c.Unlock()
	c.queue.MoveToFront(ele)
	item := ele.Value.(*cacheItem)
	item.increaseUseCount()
	return item.value, nil
}

func shouldEvictBeforeInsert(c *cache) bool {
	return c.queue.Len() >= c.size
}
