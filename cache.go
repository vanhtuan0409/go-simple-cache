// Package cache provides a dead simple implementation of in-mem cache with LRU or LFU eviction method
package cache

import (
	"container/list"
	"errors"
	"sync"
)

var (
	// ErrKeyNotFound Error when access an unknowned key in cache
	ErrKeyNotFound = errors.New("Key not found")
)

// Cache Most simplest interface for cache with 2 method Get and Set
type Cache interface {
	Set(key string, value interface{})
	Get(key string) (interface{}, error)
}

type cache struct {
	sync.Mutex
	size                   int
	queue                  *list.List
	evict                  evictFn
	shouldUpdateQueueOnGet bool
	items                  map[string]*list.Element
}

// NewLRUCache Create new cache with LRU eviction
func NewLRUCache(size int) Cache {
	return &cache{
		size:  size,
		queue: list.New(),
		items: map[string]*list.Element{},
		evict: lruEvictFn,
	}
}

// NewLFUCache Create new cache with LFU eviction
func NewLFUCache(size int) Cache {
	return &cache{
		size:  size,
		queue: list.New(),
		items: map[string]*list.Element{},
		evict: lfuEvictFn,
	}
}

// NewFIFOCache Create new cache with FIFO eviction
func NewFIFOCache(size int) Cache {
	return &cache{
		size:  size,
		queue: list.New(),
		items: map[string]*list.Element{},
		shouldUpdateQueueOnGet: true,
		evict: lruEvictFn,
	}
}

// Set Set a new cache value, will trigger eviction if cache is full
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

// Get Get a value from cache by key, return error if key is not in cache
func (c *cache) Get(key string) (interface{}, error) {
	ele, ok := c.items[key]
	if !ok {
		return nil, ErrKeyNotFound
	}
	item := ele.Value.(*cacheItem)
	item.increaseUseCount()

	if !c.shouldUpdateQueueOnGet {
		return item.value, nil
	}

	c.Lock()
	defer c.Unlock()
	c.queue.MoveToFront(ele)
	return item.value, nil
}

func shouldEvictBeforeInsert(c *cache) bool {
	return c.queue.Len() >= c.size
}
