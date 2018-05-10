package cache

import (
	"strconv"
	"testing"
)

func BenchmarkCacheWrite(b *testing.B) {
	c := NewLRUCache(100)
	for n := 0; n < b.N; n++ {
		c.Set(strconv.Itoa(n), 1)
	}
}

func BenchmarkCacheRead(b *testing.B) {
	c := NewLRUCache(100)
	for i := 0; i < 100; i++ {
		c.Set(strconv.Itoa(i), 1)
	}

	for n := 0; n < b.N; n++ {
		c.Get("1")
	}
}
