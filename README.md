# Go Simple Cache

Dead simple in-mem cache with LFU/LRU/FIFO eviction

[![GoDoc](https://godoc.org/github.com/vanhtuan0409/go-simple-cache?status.png)](https://godoc.org/github.com/vanhtuan0409/go-simple-cache)

### Installation

```
go get github.com/vanhtuan0409/go-simple-cache
```

### Usage

* Cache interface have 2 method for Get and Set value

```
type Cache interface {
	Set(key string, value interface{})
	Get(key string) (interface{}, error)
}
```

* There are 3 methods to create a cache based on eviction method:

```
func NewLRUCache(size int) Cache
func NewLFUCache(size int) Cache
func NewFIFOCache(size int) Cache
```

### Eviction

* LRU (Least Recently Used): The oldest element is the Less Recently Used (LRU) element. The last used timestamp is updated when an element is put into the cache or an element is retrieved from the cache with a get call.

```
With Cache of size 2:
Actions: Put x -> Put y -> Get x -> Get x -> Get y -> (Put z)
When z is put, x will be evict
```

* LFU (Least Frequently Used): For each get call on the element the number of hits is updated. When a put call is made for a new element (and assuming that the max limit is reached) the element with least number of hits is evicted.

```
With Cache of size 2:
Actions: Put x -> Put y -> Get x -> Get x -> Get y -> (Put z)
When z is put, y will be evict
```

* FIFO (First In First Out): Elements are evicted in the same order as they come in. When a put call is made for a new element (and assuming that the max limit is reached for the memory store) the element that was placed first (First-In) in the store is the candidate for eviction (First-Out).

```
With Cache of size 2:
Actions: Put x -> Put y -> Get x -> Get x -> Get y -> (Put z)
When z is put, x will be evict
```
